package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/golang/protobuf/ptypes/empty"
	task "github.com/tommy-sho/golang-sandbox/grpc-gateway/genproto"
)

const (
	port = "10000"
)

var tasks = []*task.Task{
	&task.Task{
		Id:   "1",
		Body: "Meet doctor",
	},
	&task.Task{
		Id:   "2",
		Body: "Watch the new movie",
	},
	&task.Task{
		Id:   "3",
		Body: "eliminate enemy",
	},
}

type ServiceImpl struct{}

func (s *ServiceImpl) GetTask(ctx context.Context, request *task.GetTaskRequest) (*task.Task, error) {
	log.Println("GetTask in gRPC server")
	for _, t := range tasks {
		if t.Id == request.Id {
			return t, nil
		}
	}
	return &task.Task{}, nil
}

func (s *ServiceImpl) GetTaskList(_ *empty.Empty, stream task.TaskManager_GetTaskListServer) error {
	log.Println("GetTaskList in gRPC server")
	for _, t := range tasks {
		if err := stream.Send(t); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	task.RegisterTaskManagerServer(s, &ServiceImpl{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("gRPC Server started: localhost%s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
