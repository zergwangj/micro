package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/zergwangj/micro/api/auth/client"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/envvar"
	"github.com/micro/go-config/source/flag"
	"github.com/micro/go-config/source/file"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"os"
	opentracinggo "github.com/opentracing/opentracing-go"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"
	microgrpc "github.com/micro/go-grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"context"
	authmicro "github.com/zergwangj/micro/api/auth/proto/micro/auth"
	"net/http"
	"github.com/zergwangj/micro/api/auth/handler"
	"google.golang.org/grpc"
	authrest "github.com/zergwangj/micro/api/auth/proto/rest/auth"
)

const serviceName = "zergwangj.api.auth"

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
		micro.WrapHandler(client.AuthWrapper(conf)),
		micro.WrapHandler(opentracing.NewHandlerWrapper(tracer)),
	)
	authmicro.RegisterAuthHandler(service.Server(), new(handler.Auth))

	// Initialise rest service
	serverAddress := conf.Get("micro", "server", "address").String(":5353")
	restfulAddress := conf.Get("restful", "address").String(":8080")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = authrest.RegisterAuthHandlerFromEndpoint(ctx, mux, serverAddress, opts)
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
