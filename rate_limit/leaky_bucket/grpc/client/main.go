package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/ratelimit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc_ratelimit "github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
	requestTime = 100
)

type T struct {
	count   int
	limiter ratelimit.Limiter
}

func (t *T) Limit() bool {
	t.count++
	fmt.Printf("count: %d\n", t.count)
	t.limiter.Take()
	return false
}

func NewRateLimiter(t int) *T {
	return &T{
		count:   0,
		limiter: ratelimit.New(t),
	}
}

func UnaryClientInterceptor(limiter grpc_ratelimit.Limiter) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if limiter.Limit() {
			return status.Errorf(codes.ResourceExhausted, "%s is rejected by grpc_ratelimit middleware, please retry later.", method)
		}

		log.Printf("method: %v", method)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithUnaryInterceptor(UnaryClientInterceptor(NewRateLimiter(1))))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	for i := 0; i < requestTime; i++ {
		if err := sendRequest(c, name); err != nil {
			log.Fatal("sendRequest: %w", err)
		}
	}
}

func sendRequest(c pb.GreeterClient, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		return fmt.Errorf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())

	return nil
}
