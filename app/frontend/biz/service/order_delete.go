package service

import (
	"context"

	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	order "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/order"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	rpcorder "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type OrderDeleteService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderDeleteService(Context context.Context, RequestContext *app.RequestContext) *OrderDeleteService {
	return &OrderDeleteService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderDeleteService) Run(req *order.OrderId) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_, err = rpc.OrderClient.DeleteOrder(h.Context, &rpcorder.DeleteOrderReq{OrderId: req.Id})
	if err != nil {
		return nil, err
	}
	return
}
