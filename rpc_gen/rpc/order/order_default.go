package order

import (
	"context"
	order "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func ListOrder(ctx context.Context, req *order.ListOrderReq, callOptions ...callopt.Option) (resp *order.ListOrderResp, err error) {
	resp, err = defaultClient.ListOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddOrder(ctx context.Context, req *order.AddOrderReq, callOptions ...callopt.Option) (resp *order.AddOrderResp, err error) {
	resp, err = defaultClient.AddOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteOrder(ctx context.Context, req *order.DeleteOrderReq, callOptions ...callopt.Option) (resp *order.DeleteOrderResp, err error) {
	resp, err = defaultClient.DeleteOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
