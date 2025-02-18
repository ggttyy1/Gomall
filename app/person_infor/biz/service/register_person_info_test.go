package service

import (
	"context"
	"testing"
	person_infor "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/person_infor"
)

func TestRegisterPersonInfo_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRegisterPersonInfoService(ctx)
	// init req and assert value

	req := &person_infor.PersonReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
