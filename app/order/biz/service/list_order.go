package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/biz-demo/gomall/app/order/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/order/biz/model"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout"
	order "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
	fmt.Println("order_list_server")
	orders, err := model.ListOrder(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, err
	}
	fmt.Println("order_list_server")
	var order_list []*order.Order
	for _, v := range orders {
		var items []*order.ProductItems
		for _, item := range v.Items {
			items = append(items, &order.ProductItems{
				ProductId: item.ProductId,
				Quantity:  item.Qty,
				Cost:      item.Cost,
			})
		}
		if v.DeletedAt.Time.IsZero() {
			order_list = append(order_list, &order.Order{
				Items:     items,
				OrderId:   v.OrderId,
				UserId:    v.UserId,
				Email:     v.Consignee.Email,
				CreatedAt: int32(v.CreatedAt.Unix()),
				Address: &checkout.Address{
					StreetAddress: v.Consignee.StreetAddress,
					City:          v.Consignee.City,
					State:         v.Consignee.State,
					Country:       v.Consignee.Country,
					ZipCode:       v.Consignee.ZipCode,
				},
				IsDelete: 0,
			})
		} else {
			order_list = append(order_list, &order.Order{
				Items:     items,
				OrderId:   v.OrderId,
				UserId:    v.UserId,
				Email:     v.Consignee.Email,
				CreatedAt: int32(v.CreatedAt.Unix()),
				Address: &checkout.Address{
					StreetAddress: v.Consignee.StreetAddress,
					City:          v.Consignee.City,
					State:         v.Consignee.State,
					Country:       v.Consignee.Country,
					ZipCode:       v.Consignee.ZipCode,
				},
				IsDelete: 1,
			})
		}

	}
	fmt.Println("order_list_server")
	return &order.ListOrderResp{OrderList: order_list}, nil
}
