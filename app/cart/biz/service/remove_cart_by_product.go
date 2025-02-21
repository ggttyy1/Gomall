package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/cart/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/cart/biz/model"
	cart "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
)

type RemoveCartByProductService struct {
	ctx context.Context
} // NewRemoveCartByProductService new RemoveCartByProductService
func NewRemoveCartByProductService(ctx context.Context) *RemoveCartByProductService {
	return &RemoveCartByProductService{ctx: ctx}
}

// Run create note info
func (s *RemoveCartByProductService) Run(req *cart.RemoveCartByProductReq) (resp *cart.RemoveCartByProductResp, err error) {
	// Finish your business logic.
	err = model.DeleteUserProduct(s.ctx, mysql.DB, req.UserId, req.ProductIds)
	if err != nil {
		return nil, err
	}
	return
}
