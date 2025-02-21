package service

import (
	"context"
	"testing"
	checkout "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout"
)

func TestCheckoutProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCheckoutProductService(ctx)
	// init req and assert value

	req := &checkout.CheckoutProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
