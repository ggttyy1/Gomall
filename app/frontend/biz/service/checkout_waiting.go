package service

import (
	"context"
	"fmt"

	checkout "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/checkout"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	rpccheckout "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	fmt.Println(req.SelectedPerson)
	userId := frontendutils.GetUserIdfromCtx(h.Context)
	_, err = rpc.CheckoutClient.CheckoutProduct(h.Context, &rpccheckout.CheckoutProductReq{
		UserId:        uint32(userId),
		PersonInforId: req.SelectedPerson,
		ProductIds:    req.SelectedProduct,
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title":    "waiting",
		"redirect": "/checkout/result",
	}, nil
}
