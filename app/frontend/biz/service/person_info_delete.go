package service

import (
	"context"

	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	person_info "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/person_info"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/person_infor"
	"github.com/cloudwego/hertz/pkg/app"
)

type PersonInfoDeleteService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPersonInfoDeleteService(Context context.Context, RequestContext *app.RequestContext) *PersonInfoDeleteService {
	return &PersonInfoDeleteService{RequestContext: RequestContext, Context: Context}
}

func (h *PersonInfoDeleteService) Run(req *person_info.PersonInfoDeleteReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_, err = rpc.PersonInfoClient.DeletePersonInfo(h.Context, &person_infor.DeletePersonInfoReq{
		PersonInfoId: req.PersonId,
	})
	if err != nil {
		return nil, err
	}
	return
}
