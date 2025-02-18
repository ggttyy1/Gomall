package product

import (
	"context"
	product "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() productcatalogservice.Client
	Service() string
	ListProducts(ctx context.Context, Req *product.ListProductsRequest, callOptions ...callopt.Option) (r *product.ListProductsResponse, err error)
	GetProduct(ctx context.Context, Req *product.GetProductRequest, callOptions ...callopt.Option) (r *product.GetProductResponse, err error)
	SearchProducts(ctx context.Context, Req *product.SearchProductsRequest, callOptions ...callopt.Option) (r *product.SearchProductsResponse, err error)
	UploadProduct(ctx context.Context, Req *product.UploadProductRequest, callOptions ...callopt.Option) (r *product.UploadProductResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := productcatalogservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient productcatalogservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() productcatalogservice.Client {
	return c.kitexClient
}

func (c *clientImpl) ListProducts(ctx context.Context, Req *product.ListProductsRequest, callOptions ...callopt.Option) (r *product.ListProductsResponse, err error) {
	return c.kitexClient.ListProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) GetProduct(ctx context.Context, Req *product.GetProductRequest, callOptions ...callopt.Option) (r *product.GetProductResponse, err error) {
	return c.kitexClient.GetProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) SearchProducts(ctx context.Context, Req *product.SearchProductsRequest, callOptions ...callopt.Option) (r *product.SearchProductsResponse, err error) {
	return c.kitexClient.SearchProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) UploadProduct(ctx context.Context, Req *product.UploadProductRequest, callOptions ...callopt.Option) (r *product.UploadProductResponse, err error) {
	return c.kitexClient.UploadProduct(ctx, Req, callOptions...)
}
