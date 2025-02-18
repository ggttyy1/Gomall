package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/product/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/product/biz/model"
	product "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
)

type UploadProductService struct {
	ctx context.Context
} // NewUploadProductService new UploadProductService
func NewUploadProductService(ctx context.Context) *UploadProductService {
	return &UploadProductService{ctx: ctx}
}

// Run create note info
func (s *UploadProductService) Run(req *product.UploadProductRequest) (resp *product.UploadProductResponse, err error) {
	// 开始一个数据库事务
	tx := mysql.DB.Begin()

	// 如果事务开始失败，返回错误
	if tx.Error != nil {
		return nil, tx.Error
	}

	ProCate, err := model.SearchCategory(s.ctx, tx, req.Product.ProductCategory)
	if err != nil {
		tx.Rollback() // 如果创建分类失败，回滚事务
		return nil, err
	}

	// 创建商品
	p := model.Product{
		Name:        req.Product.Name,
		Description: req.Product.Description,
		Price:       req.Product.Price,
		Picture:     req.Product.Picture,
		Categories:  []model.Category{*ProCate},
	}
	err = model.CreateProduct(s.ctx, tx, p)
	if err != nil {
		tx.Rollback() // 如果创建商品失败，回滚事务
		return nil, err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return
}
