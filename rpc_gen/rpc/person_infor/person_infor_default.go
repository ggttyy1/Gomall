package person_infor

import (
	"context"
	person_infor "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/person_infor"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func RegisterPersonInfo(ctx context.Context, req *person_infor.PersonReq, callOptions ...callopt.Option) (resp *person_infor.PersonResp, err error) {
	resp, err = defaultClient.RegisterPersonInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "RegisterPersonInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetPersonInfoList(ctx context.Context, req *person_infor.GetPersonInfoListReq, callOptions ...callopt.Option) (resp *person_infor.GetPersonInfoListResp, err error) {
	resp, err = defaultClient.GetPersonInfoList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetPersonInfoList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetPersonInfo(ctx context.Context, req *person_infor.GetPersonInfoReq, callOptions ...callopt.Option) (resp *person_infor.GetPersonInfoResp, err error) {
	resp, err = defaultClient.GetPersonInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetPersonInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdatePersonInfo(ctx context.Context, req *person_infor.UpdatePersonInfoReq, callOptions ...callopt.Option) (resp *person_infor.UpdatePersonInfoResp, err error) {
	resp, err = defaultClient.UpdatePersonInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdatePersonInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeletePersonInfo(ctx context.Context, req *person_infor.DeletePersonInfoReq, callOptions ...callopt.Option) (resp *person_infor.DeletePersonInfoResp, err error) {
	resp, err = defaultClient.DeletePersonInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeletePersonInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
