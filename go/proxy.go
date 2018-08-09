package main

import (
	"flag"
	"net/http"

	"golang.org/x/net/context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"

	pb "test-grpc/example"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:9090", "endpoint of YourService")
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	customMarshaller := &runtime.JSONPb{
		OrigName:     true,
		EmitDefaults: true, // disable 'omitempty'
	}
	muxOpt := runtime.WithMarshalerOption(runtime.MIMEWildcard, customMarshaller)
	gw := runtime.NewServeMux(muxOpt)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterYourServiceHandlerFromEndpoint(ctx, gw, *echoEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gw)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
