package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/order/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/order/biz/model"
	order "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
)

type DeleteOrderService struct {
	ctx context.Context
} // NewDeleteOrderService new DeleteOrderService
func NewDeleteOrderService(ctx context.Context) *DeleteOrderService {
	return &DeleteOrderService{ctx: ctx}
}

// Run create note info
func (s *DeleteOrderService) Run(req *order.DeleteOrderReq) (resp *order.DeleteOrderResp, err error) {
	// Finish your business logic.
	err = model.DeleteOrder(s.ctx, mysql.DB, req.OrderId)
	if err != nil {
		return nil, err
	}
	return
}
