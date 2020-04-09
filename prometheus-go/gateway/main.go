package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	echoPrometheus "github.com/globocom/echo-prometheus"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"

	"github.com/tommy-sho/k8s-grpc-health-check/proto"
)

const (
	port     = ":8080"
	promAddr = ":9090"
)

var backendPort string

func init() {
	backendPort = os.Getenv("ENDPOINT")
	if backendPort == "" {
		backendPort = ":50001"
	}
}

type Request struct {
	Name string `json:"name" form:"name" query:"name"`
}

type Response struct {
	Message  string `json:"message"`
	DateTime string `json:"datetime"`
}

func main() {
	dialOpts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
	}

	bConn, err := grpc.Dial(backendPort, dialOpts...)
	if err != nil {
		panic(fmt.Errorf("failed to connect with backend server error : %v ", err))
	}

	fmt.Printf("%s:%s", os.Getenv("MY_POD_IP"), os.Getenv("BACKEND_PORT"))

	bClient := proto.NewBackendServerClient(bConn)

	var configMetrics = echoPrometheus.NewConfig()
	configMetrics.Namespace = "namespace"
	configMetrics.Buckets = []float64{
		0.0005, // 0.5ms
		0.001,  // 1ms
		0.005,  // 5ms
		0.01,   // 10ms
		0.05,   // 50ms
		0.1,    // 100ms
		0.5,    // 500ms
		1,      // 1s
		2,      // 2s
	}
	e := echo.New()
	e.GET("/greeting", Greeting(bClient))
	e.GET("/healthz", Pong)
	e.Use(echoPrometheus.MetricsMiddlewareWithConfig(configMetrics))
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	mux := http.NewServeMux()
	// Enable histogram
	grpc_prometheus.EnableClientHandlingTimeHistogram()
	mux.Handle("/metrics", promhttp.Handler())
	go func() {
		fmt.Println("Prometheus metrics bind address:", promAddr)
		log.Fatal(http.ListenAndServe(promAddr, mux))
	}()

	go func() {
		<-stopChan
		if err := e.Close(); err != nil {
			log.Print("Failed to stop server")
		}
	}()

	errors := make(chan error)
	go func() {
		errors <- e.Start(port)
	}()

	if err := <-errors; err != nil {
		log.Fatal("Failed to server gRPC server", err)
	}
}

func Greeting(client proto.BackendServerClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := new(Request)
		if err := c.Bind(r); err != nil {
			return err
		}

		ctx := context.Background()
		req := &proto.MessageRequest{
			Name: r.Name,
		}
		m, err := client.Message(ctx, req)
		if err != nil {
			log.Printf("failed to access to backend service")
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		res := Response{
			Message:  m.Message,
			DateTime: time.Unix(m.Datetime, 0).String(),
		}
		return c.JSON(http.StatusOK, res)
	}
}

func Pong(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
