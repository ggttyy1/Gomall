package service

import (
	"context"
	"testing"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/biz/dal/mysql"
	user "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestRegister_Run(t *testing.T) {
	_ = godotenv.Load()
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
