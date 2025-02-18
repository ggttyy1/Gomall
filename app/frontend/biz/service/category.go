package service

import (
	"context"
	"fmt"

	category "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/category"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryRequest) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	p, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsRequest{CatagoryName: req.Category})
	fmt.Println("收到个数：", len(p.Products))
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title": "Category",
		"item":  p.Products,
	}, nil
}
