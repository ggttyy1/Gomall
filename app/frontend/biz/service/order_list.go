package service

import (
	"context"
	"time"

	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/types"
	frontendUtil "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userId := frontendUtil.GetUserIdfromCtx(h.Context)
	orderResp, err := rpc.OrderClient.ListOrder(h.Context, &order.ListOrderReq{
		UserId: uint32(userId),
	})
	if err != nil {
		return nil, err
	}
	var resItems []types.Order
	for _, v := range orderResp.OrderList {
		var total float32
		var Items []types.OrderItem
		for _, item := range v.Items {
			total += item.Cost
			productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductRequest{Id: item.ProductId})
			if err != nil {
				return nil, err
			}
			Items = append(Items, types.OrderItem{
				ProductName: productResp.Product.Name,
				Qty:         item.Quantity,
				Picture:     productResp.Product.Picture,
				Cost:        item.Cost,
			})
		}
		created := time.Unix(int64(v.CreatedAt), 0)
		resItems = append(resItems, types.Order{
			OrderId:     v.OrderId,
			CreatedDate: created.Format("2006-01-02 15:04:05"),
			Items:       Items,
			Cost:        total,
			Deleted:     v.IsDelete,
		})

	}
	return utils.H{
		"orders": resItems,
		"Title":  "order",
	}, nil
}
