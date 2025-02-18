package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/product/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/product/biz/model"
	product "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsRequest) (resp *product.SearchProductsResponse, err error) {
	// Finish your business logic.
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	p, err := productQuery.SearchProduct(req.Query)
	if err != nil {
		return nil, err
	}
	resp = &product.SearchProductsResponse{}
	for _, v := range p {
		resp.Results = append(resp.Results, &product.Product{
			Id:          uint32(v.ID),
			Description: v.Description,
			Price:       v.Price,
			Picture:     v.Picture,
			Name:        v.Name,
		})
	}
	return
}
