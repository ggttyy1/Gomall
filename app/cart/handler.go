package main

import (
	"context"
	cart "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/biz-demo/gomall/app/cart/biz/service"
)

// CartrSeviceImpl implements the last service interface defined in the IDL.
type CartrSeviceImpl struct{}

// AddItem implements the CartrSeviceImpl interface.
func (s *CartrSeviceImpl) AddItem(ctx context.Context, req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	resp, err = service.NewAddItemService(ctx).Run(req)

	return resp, err
}

// GetCart implements the CartrSeviceImpl interface.
func (s *CartrSeviceImpl) GetCart(ctx context.Context, req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	resp, err = service.NewGetCartService(ctx).Run(req)

	return resp, err
}

// EmptyCart implements the CartrSeviceImpl interface.
func (s *CartrSeviceImpl) EmptyCart(ctx context.Context, req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	resp, err = service.NewEmptyCartService(ctx).Run(req)

	return resp, err
}

// GetCartByProduct implements the CartrSeviceImpl interface.
func (s *CartrSeviceImpl) GetCartByProduct(ctx context.Context, req *cart.GetCartByProductReq) (resp *cart.GetCartByProductResp, err error) {
	resp, err = service.NewGetCartByProductService(ctx).Run(req)

	return resp, err
}

// RemoveCartByProduct implements the CartrSeviceImpl interface.
func (s *CartrSeviceImpl) RemoveCartByProduct(ctx context.Context, req *cart.RemoveCartByProductReq) (resp *cart.RemoveCartByProductResp, err error) {
	resp, err = service.NewRemoveCartByProductService(ctx).Run(req)

	return resp, err
}
