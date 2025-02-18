package service

import (
	"context"
	"testing"

	"github.com/cloudwego/biz-demo/gomall/app/cart/rpc"
	cart "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
)

func TestGetCart_Run(t *testing.T) {
	rpc.TestInit()
	ctx := context.Background()

	s := NewGetCartService(ctx)
	// init req and assert value

	req := &cart.GetCartReq{UserId: 1}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
