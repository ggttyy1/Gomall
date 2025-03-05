package middleware

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/biz/utils"
	frontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

// GlobalAuth 函数返回一个 app.HandlerFunc 类型的函数，用于全局身份验证
func GlobalAuth() app.HandlerFunc {
	// 返回一个匿名函数，该函数接收两个参数：context.Context 类型的 c 和 *app.RequestContext 类型的 ctx
	return func(c context.Context, ctx *app.RequestContext) {
		// TODO: implement
		// 获取默认的 session 对象
		s := sessions.Default(ctx)
		// 将 session 中的 user_id 存入 context 中
		c = context.WithValue(c, frontendUtils.SessionUserId, s.Get("user_id"))
		// 调用 ctx.Next 函数，继续处理请求
		ctx.Next(c)
	}
}

func Auth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// TODO: implement
		// s := sessions.Default(ctx)
		// userId := s.Get("user_id")

		UserID := c.Value(frontendUtils.JWTUserId)
		if UserID.(int) <= 0 {
			ctx.Redirect(302, []byte("/sign-in?next="+ctx.FullPath()))
			ctx.Abort()
			return
		}
		ctx.Next(c)
	}

}
func JWTGlobalAuth() app.HandlerFunc {
	// 返回一个匿名函数，该函数接收两个参数：context.Context 类型的 c 和 *app.RequestContext 类型的 ctx
	return func(c context.Context, ctx *app.RequestContext) {
		_, data, _ := utils.ParseJWT(string(ctx.Cookie("token")))
		c = context.WithValue(c, frontendUtils.JWTUserId, data.UserID)
		c = context.WithValue(c, frontendUtils.JWTRole, data.Role)
		ctx.Next(c)
	}
}
