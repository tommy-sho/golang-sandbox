package handler

import (
	// ...

	"context"
	"net/http"
	"strconv"

	"github.com/tommy-sho/golang-sandbox/web/payjp-vue/backend-api/database"

	"github.com/tommy-sho/golang-sandbox/web/payjp-vue/backend-api/domain"
	gpay "github.com/tommy-sho/golang-sandbox/web/payjp-vue/proto"

	"google.golang.org/grpc"
)

var addr = "localhost:50051"

// Charge exec payment-service charge
func Charge(c Context) {
	//パラメータや body をうけとる
	t := domain.Payment{}
	c.Bind(&t)
	identifer, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	// id から item情報取得
	res, err := database.SelectItem(int64(identifer))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	// gRPC サーバーに送る Request を作成
	greq := &gpay.PayRequest{
		Id:          int64(identifer),
		Token:       t.Token,
		Amount:      res.Amount,
		Name:        res.Name,
		Description: res.Description,
	}

	//IPアドレス(ここではlocalhost)とポート番号(ここでは50051)を指定して、サーバーと接続する
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		c.JSON(http.StatusForbidden, err)
	}
	defer conn.Close()
	client := gpay.NewPayManagerClient(conn)

	// gRPCマイクロサービスの支払い処理関数を叩く
	gres, err := client.Charge(context.Background(), greq)
	if err != nil {
		c.JSON(http.StatusForbidden, err)
		return
	}
	c.JSON(http.StatusOK, gres)
}
