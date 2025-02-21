package checkout

import (
	"context"
	checkout "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Checkout(ctx context.Context, req *checkout.CheckoutReq, callOptions ...callopt.Option) (resp *checkout.CheckoutResp, err error) {
	resp, err = defaultClient.Checkout(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Checkout call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CheckoutProduct(ctx context.Context, req *checkout.CheckoutProductReq, callOptions ...callopt.Option) (resp *checkout.CheckoutProductResp, err error) {
	resp, err = defaultClient.CheckoutProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CheckoutProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
