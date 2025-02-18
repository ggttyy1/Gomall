package main

import (
	"context"
	order "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/biz-demo/gomall/app/order/biz/service"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	resp, err = service.NewListOrderService(ctx).Run(req)

	return resp, err
}

// AddOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) AddOrder(ctx context.Context, req *order.AddOrderReq) (resp *order.AddOrderResp, err error) {
	resp, err = service.NewAddOrderService(ctx).Run(req)

	return resp, err
}
