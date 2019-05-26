package main

import(
	"../proto"
	"context"
)

const (
	port = ":50001"
)

type paymentServer struct {}

func (t *paymentServer)Charge( ctx context.Context, req *paymentservice.PayRequest) (*paymentservice.PayResponse, error){
	pay := payjp
	return &paymentservice.PayResponse{}, nil
}




func main(){

}