package service

import (
	"context"
	"time"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/biz/utils"
	auth "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {
	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return "", err
	}

	tokenString, err := utils.GenerateJWT(int(resp.UserId), "user")
	if err != nil {
		return "", err
	}
	// 设置 cookie
	expireCookie := time.Now().Add(24 * time.Hour)         // 设置 cookie 过期时间为 1 天
	maxage := int(expireCookie.Unix() - time.Now().Unix()) // 计算 maxage

	h.RequestContext.SetCookie("token", tokenString, maxage, "/", "", protocol.CookieSameSiteDefaultMode, true, true)

	if err != nil {
		return "", err
	}

	if req.Next != "" {
		redirect = req.Next
	} else {
		redirect = "/"
	}
	return
}
