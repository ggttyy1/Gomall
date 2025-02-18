package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/person_infor/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/person_infor/biz/model"
	person_infor "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/person_infor"
)

type DeletePersonInfoService struct {
	ctx context.Context
} // NewDeletePersonInfoService new DeletePersonInfoService
func NewDeletePersonInfoService(ctx context.Context) *DeletePersonInfoService {
	return &DeletePersonInfoService{ctx: ctx}
}

// Run create note info
func (s *DeletePersonInfoService) Run(req *person_infor.DeletePersonInfoReq) (resp *person_infor.DeletePersonInfoResp, err error) {
	// Finish your business logic.
	err = model.Delete(s.ctx, mysql.DB, req.PersonInfoId)
	if err != nil {
		return nil, err
	}
	return
}
