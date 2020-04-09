package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/tommy-sho/k8s-grpc-health-check/proto"
)

const (
	port     = ":50001"
	promAddr = ":9090"
)

func main() {
	b := &BackendServer{}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	proto.RegisterBackendServerServer(server, b)
	grpc_prometheus.Register(server)

	reflection.Register(server)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	mux := http.NewServeMux()
	// Enable histogram
	grpc_prometheus.EnableHandlingTimeHistogram()
	mux.Handle("/metrics", promhttp.Handler())
	go func() {
		fmt.Println("Prometheus metrics bind address", promAddr)
		log.Fatal(http.ListenAndServe(promAddr, mux))
	}()

	go func() {
		<-stopChan
		gracefulStopChan := make(chan bool, 1)
		go func() {
			server.GracefulStop()
			gracefulStopChan <- true
		}()
		t := time.NewTimer(10 * time.Second)
		select {
		case <-gracefulStopChan:
			log.Print("Success graceful stop")
		case <-t.C:
			server.Stop()
		}
	}()

	errors := make(chan error)
	go func() {
		errors <- server.Serve(lis)
	}()

	if err := <-errors; err != nil {
		log.Fatal("Failed to server gRPC server", err)
	}

}

type BackendServer struct{}

func (b *BackendServer) Message(ctx context.Context, req *proto.MessageRequest) (*proto.MessageResponse, error) {
	message := fmt.Sprintf("Hey! %s, Nice to meet you!!", req.Name)
	currentTime := time.Now()

	log.Println(message)
	res := &proto.MessageResponse{
		Message:  message,
		Datetime: currentTime.Unix(),
	}
	return res, nil
}
