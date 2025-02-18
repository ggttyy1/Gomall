package main

import (
	"context"
	person_infor "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/person_infor"
	"github.com/cloudwego/biz-demo/gomall/app/person_infor/biz/service"
)

// PersonInforImpl implements the last service interface defined in the IDL.
type PersonInforImpl struct{}

// RegisterPersonInfo implements the PersonInforImpl interface.
func (s *PersonInforImpl) RegisterPersonInfo(ctx context.Context, req *person_infor.PersonReq) (resp *person_infor.PersonResp, err error) {
	resp, err = service.NewRegisterPersonInfoService(ctx).Run(req)

	return resp, err
}

// GetPersonInfo implements the PersonInforImpl interface.
func (s *PersonInforImpl) GetPersonInfo(ctx context.Context, req *person_infor.GetPersonInfoReq) (resp *person_infor.GetPersonInfoResp, err error) {
	resp, err = service.NewGetPersonInfoService(ctx).Run(req)

	return resp, err
}

// GetPersonInfoList implements the PersonInforImpl interface.
func (s *PersonInforImpl) GetPersonInfoList(ctx context.Context, req *person_infor.GetPersonInfoListReq) (resp *person_infor.GetPersonInfoListResp, err error) {
	resp, err = service.NewGetPersonInfoListService(ctx).Run(req)

	return resp, err
}

// UpdatePersonInfo implements the PersonInforImpl interface.
func (s *PersonInforImpl) UpdatePersonInfo(ctx context.Context, req *person_infor.UpdatePersonInfoReq) (resp *person_infor.UpdatePersonInfoResp, err error) {
	resp, err = service.NewUpdatePersonInfoService(ctx).Run(req)

	return resp, err
}

// DeletePersonInfo implements the PersonInforImpl interface.
func (s *PersonInforImpl) DeletePersonInfo(ctx context.Context, req *person_infor.DeletePersonInfoReq) (resp *person_infor.DeletePersonInfoResp, err error) {
	resp, err = service.NewDeletePersonInfoService(ctx).Run(req)

	return resp, err
}
