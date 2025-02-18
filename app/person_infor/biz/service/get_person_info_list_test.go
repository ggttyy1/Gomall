package service

import (
	"context"
	"testing"
	person_infor "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/person_infor"
)

func TestGetPersonInfoList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetPersonInfoListService(ctx)
	// init req and assert value

	req := &person_infor.GetPersonInfoListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
