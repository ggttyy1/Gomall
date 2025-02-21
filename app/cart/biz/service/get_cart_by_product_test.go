package service

import (
	"context"
	"testing"
	cart "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
)

func TestGetCartByProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetCartByProductService(ctx)
	// init req and assert value

	req := &cart.GetCartByProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
