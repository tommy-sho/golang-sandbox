package main

import (
	"context"
	"net"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/payjp/payjp-go/v1"
	ps "github.com/tommy-sho/golang-sandbox/web/payjp-vue/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50001"
)

type paymentServer struct{}

func (t *paymentServer) Charge(ctx context.Context, req *ps.PayRequest) (*ps.PayResponse, error) {
	pay := payjp.New(os.Getenv("PAYJP_TEST_SECRET_KEY"), nil)

	charge, err := pay.Charge.Create(int(req.Amount), payjp.Charge{
		Currency:    "jpy",
		CardToken:   req.Token,
		Capture:     true,
		Description: req.Name + ":" + req.Description,
	})
	if err != nil {
		return nil, err
	}

	res := &ps.PayResponse{
		Paid:     charge.Paid,
		Captured: charge.Captured,
		Amount:   int64(charge.Amount),
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen :%v", err)
	}
	s := grpc.NewServer()
	ps.RegisterPayManagerServer(s, &paymentServer{})
	reflection.Register(s)
	log.Printf("start listening grpc server: localhost%v", port)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
