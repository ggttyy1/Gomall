package service

import (
	"context"
	"testing"
	cart "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
)

func TestRemoveCartByProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRemoveCartByProductService(ctx)
	// init req and assert value

	req := &cart.RemoveCartByProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
