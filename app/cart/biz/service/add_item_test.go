package service

import (
	"context"
	"testing"

	"github.com/cloudwego/biz-demo/gomall/app/cart/biz/dal"
	"github.com/cloudwego/biz-demo/gomall/app/cart/rpc"
	cart "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
)

func TestAddItem_Run(t *testing.T) {
	rpc.TestInit()
	dal.TestInit()
	ctx := context.Background()
	s := NewAddItemService(ctx)
	// init req and assert value

	req := &cart.AddItemReq{UserId: 1, Items: &cart.CartItem{ProductId: 1, Quantity: 10}}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

func TestRpcRun(t *testing.T) {
	// Finish your business logic.
	rpc.TestInit()
	ctx := context.Background()
	productResp, err := rpc.ProductClient.GetProduct(ctx, &product.GetProductRequest{Id: 1})

	if err != nil {
		t.Logf("err: %v", err)
	}
	t.Logf(productResp.Product.Description)
}
