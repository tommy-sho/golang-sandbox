package main

import (
	"context"
	"flag"
	"net/http"

	gw "github.com/tommy-sho/golang-sandbox/grpc-gateway/proto"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	tasklistEndpoint = flag.String("tasklist_endpoint", "localhost:10000", "endpoint of YourService")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err :=
		//RegisterTaskManagerHandlerFromEndpoint(ctx, mux, *tasklistEndpoint, opts)

	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
