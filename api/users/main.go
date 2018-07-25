package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/zergwangj/micro/api/users/handler"
	"github.com/zergwangj/micro/api/users/client"
	usersmicro "github.com/zergwangj/micro/api/users/proto/micro/users"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/envvar"
	"github.com/micro/go-config/source/flag"
	"github.com/micro/go-config/source/file"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"os"
	opentracinggo "github.com/opentracing/opentracing-go"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"
	microgrpc "github.com/micro/go-grpc"
	authwrapper "github.com/zergwangj/micro/srv/auth/wrapper"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"context"
	"google.golang.org/grpc"
	"net/http"
	usersrest "github.com/zergwangj/micro/api/users/proto/rest/users"
)

const serviceName = "zergwangj.api.users"

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

	// Initialise grpc service
	service := microgrpc.NewService(
		micro.Name(serviceName),
		micro.Version("latest"),
	)
	service.Init(
		micro.WrapHandler(authwrapper.NewHandlerWrapper()),
		micro.WrapHandler(client.UsersWrapper(conf)),
		micro.WrapHandler(opentracing.NewHandlerWrapper(tracer)),
	)
	usersmicro.RegisterUsersHandler(service.Server(), new(handler.Users))

	// Initialise rest service
	serverAddress := conf.Get("micro", "server", "address").String(":5353")
	restfulAddress := conf.Get("restful", "address").String(":8080")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = usersrest.RegisterUsersHandlerFromEndpoint(ctx, mux, serverAddress, opts)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Run service
	go http.ListenAndServe(restfulAddress, mux)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
