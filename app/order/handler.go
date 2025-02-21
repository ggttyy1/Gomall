package main

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/order/biz/service"
	order "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
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

// DeleteOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) DeleteOrder(ctx context.Context, req *order.DeleteOrderReq) (resp *order.DeleteOrderResp, err error) {
	resp, err = service.NewDeleteOrderService(ctx).Run(req)

	return resp, err
}
