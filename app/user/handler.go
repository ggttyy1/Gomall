package main

import (
	"context"
	user "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/biz-demo/gomall/app/user/biz/service"
)

// UserSeviceImpl implements the last service interface defined in the IDL.
type UserSeviceImpl struct{}

// Register implements the UserSeviceImpl interface.
func (s *UserSeviceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// Login implements the UserSeviceImpl interface.
func (s *UserSeviceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}
