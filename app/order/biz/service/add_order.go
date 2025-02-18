package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/order/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/order/biz/model"
	order "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"github.com/google/uuid"
)

type AddOrderService struct {
	ctx context.Context
} // NewAddOrderService new AddOrderService
func NewAddOrderService(ctx context.Context) *AddOrderService {
	return &AddOrderService{ctx: ctx}
}

// Run create note info
func (s *AddOrderService) Run(req *order.AddOrderReq) (resp *order.AddOrderResp, err error) {
	// Finish your business logic.
	var items []model.OrderItems
	for _, v := range req.Items {
		items = append(items, model.OrderItems{
			ProductId: v.ProductId,
			Qty:       v.Quantity,
			Cost:      v.Cost,
		})

	}
	orderId, _ := uuid.NewUUID()
	orderitem := &model.Order{
		UserId:  req.UserId,
		Items:   items,
		OrderId: orderId.String(),
		Consignee: model.Consignee{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			ZipCode:       req.Address.ZipCode,
			Country:       req.Address.Country,
			Email:         req.Email,
		},
	}
	err = model.CreateOrder(s.ctx, mysql.DB, orderitem)
	if err != nil {
		return nil, err
	}
	return &order.AddOrderResp{OrderId: orderId.String()}, nil
}
