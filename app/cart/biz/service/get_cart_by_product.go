package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/cart/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/cart/biz/model"
	cart "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
)

type GetCartByProductService struct {
	ctx context.Context
} // NewGetCartByProductService new GetCartByProductService
func NewGetCartByProductService(ctx context.Context) *GetCartByProductService {
	return &GetCartByProductService{ctx: ctx}
}

// Run create note info
func (s *GetCartByProductService) Run(req *cart.GetCartByProductReq) (resp *cart.GetCartByProductResp, err error) {
	// Finish your business logic.
	cartItems, err := model.SearchUserProduct(s.ctx, mysql.DB, req.UserId, req.ProductIds)
	if err != nil {
		return nil, err
	}
	items := []*cart.CartItem{}
	for _, item := range cartItems {
		items = append(items, &cart.CartItem{
			Quantity:  item.Quantity,
			ProductId: item.ProductID,
		})
	}
	return &cart.GetCartByProductResp{
		Items: items,
	}, nil
}
