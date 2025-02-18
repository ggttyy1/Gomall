package service

import (
	"context"
	"strconv"

	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	cartResp, err := rpc.CartClient.GetCart(h.Context, &cart.GetCartReq{
		UserId: uint32(frontendUtils.GetUserIdfromCtx(h.Context)),
	})
	if err != nil {
		return nil, err
	}
	var items []map[string]any
	var total float64
	for _, v := range cartResp.Items {
		ProductItem, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductRequest{Id: v.ProductId})
		if err != nil {
			return nil, err
		}
		items = append(items, map[string]any{
			"Name":        ProductItem.Product.Name,
			"Description": ProductItem.Product.Description,
			"Price":       ProductItem.Product.Price,
			"Picture":     ProductItem.Product.Picture,
			"Quantity":    v.Quantity,
		})
		total += float64(ProductItem.Product.Price) * float64(v.Quantity)

	}
	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
