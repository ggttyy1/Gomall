package main

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/product/biz/service"
	product "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsRequest) (resp *product.ListProductsResponse, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductRequest) (resp *product.GetProductResponse, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsRequest) (resp *product.SearchProductsResponse, err error) {
	resp, err = service.NewSearchProductsService(ctx).Run(req)

	return resp, err
}

// UploadProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) UploadProduct(ctx context.Context, req *product.UploadProductRequest) (resp *product.UploadProductResponse, err error) {
	resp, err = service.NewUploadProductService(ctx).Run(req)

	return resp, err
}
