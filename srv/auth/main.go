package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/zergwangj/micro/srv/auth/handler"
	auth "github.com/zergwangj/micro/srv/auth/proto/auth"
	"github.com/zergwangj/micro/srv/users/models"
	"github.com/muesli/cache2go"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/envvar"
	"github.com/micro/go-config/source/flag"
	"github.com/micro/go-config/source/file"
	jujuratelimit "github.com/juju/ratelimit"
	"github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"os"
	opentracinggo "github.com/opentracing/opentracing-go"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/micro/go-grpc"
	"github.com/micro/go-plugins/broker/kafka"
	broker2 "github.com/micro/go-micro/broker"
)

const serviceName = "zergwangj.srv.auth"

func main() {
	// New Config
	conf := config.NewConfig()
	conf.Load(
		// base config from env
		envvar.NewSource(),
		// override env with flags
		flag.NewSource(),
		// override flags with file
		file.NewSource(
			file.WithPath("config.yaml"),
		),
	)

	// New zipkin trace
	zipkinAddress := conf.Get("zipkin", "address").String(":9411")
	zipkinUrl := "http://" + zipkinAddress + "/api/v1/spans"
	hostname, _ := os.Hostname()
	collector, err := zipkin.NewHTTPCollector(zipkinUrl)
	if err != nil {
		log.Fatal(err)
		return
	}
	tracer, err := zipkin.NewTracer(
		zipkin.NewRecorder(collector, false, hostname, serviceName),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	opentracinggo.InitGlobalTracer(tracer)

	// New Service
	kafkaAddress := conf.Get("kafka", "address").String(":9092")
	broker := kafka.NewBroker(broker2.Addrs(kafkaAddress))
	service := grpc.NewService(
		micro.Name(serviceName),
		micro.Version("latest"),
		micro.Broker(broker),
	)

	// Initialise service
	limit := conf.Get("ratelimit").Int(100)
	b := jujuratelimit.NewBucketWithRate(float64(limit), int64(limit))
	service.Init(
		micro.WrapHandler(ratelimit.NewHandlerWrapper(b, false)),
		micro.WrapHandler(opentracing.NewHandlerWrapper(tracer)),
	)

	// Register Handler
	dbAddress := conf.Get("db", "address").String(":3306")
	dbUser := conf.Get("db", "user").String("root")
	dbPassword := conf.Get("db", "password").String("root")
	dataSourceName := dbUser + ":" + dbPassword + "@tcp(" + dbAddress + ")/example"
	db, err := models.NewDB(dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	cache := cache2go.Cache("sessions")
	auth.RegisterAuthHandler(service.Server(), &handler.Auth{
		Db: 	db,
		Cache:	cache,
		Conf:	conf,
	})

	//// Register Struct as Subscriber
	//micro.RegisterSubscriber("zergwangj.srv.auth", service.Server(), new(subscriber.Example))
	//
	//// Register Function as Subscriber
	//micro.RegisterSubscriber("zergwangj.srv.auth", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
