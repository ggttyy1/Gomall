package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/biz-demo/gomall/app/product/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/product/biz/model"
	product "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsRequest) (resp *product.ListProductsResponse, err error) {
	// Finish your business logic.
	categoryQuery := model.NewCateGoryQuery(s.ctx, mysql.DB)
	fmt.Println("categoryQuery:", req.CatagoryName)
	categories, err := categoryQuery.GetProductsByCategoryName(req.CatagoryName)
	fmt.Print("search cates:", len(categories), len(categories[0].Products))
	fmt.Println("")
	if err != nil {
		return nil, err
	}
	resp = &product.ListProductsResponse{}
	for _, v1 := range categories {
		for _, v := range v1.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(v.ID),
				Picture:     v.Picture,
				Name:        v.Name,
				Price:       v.Price,
				Description: v.Description,
			})
		}
	}
	return
}
