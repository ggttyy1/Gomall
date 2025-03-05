package service

import (
	"context"
	"time"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/biz/utils"
	auth "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/auth"
	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	UserResp, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	})
	if err != nil {
		return nil, err
	}

	tokenString, err := utils.GenerateJWT(int(UserResp.UserId), "user")
	if err != nil {
		return resp, err
	}
	// 设置 cookie
	expireCookie := time.Now().Add(24 * time.Hour)         // 设置 cookie 过期时间为 1 天
	maxage := int(expireCookie.Unix() - time.Now().Unix()) // 计算 maxage

	h.RequestContext.SetCookie("token", tokenString, maxage, "/", "", protocol.CookieSameSiteDefaultMode, true, true)
	return

}
