package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/person_infor/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/person_infor/biz/model"
	person_infor "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/person_infor"
)

type UpdatePersonInfoService struct {
	ctx context.Context
} // NewUpdatePersonInfoService new UpdatePersonInfoService
func NewUpdatePersonInfoService(ctx context.Context) *UpdatePersonInfoService {
	return &UpdatePersonInfoService{ctx: ctx}
}

// Run create note info
func (s *UpdatePersonInfoService) Run(req *person_infor.UpdatePersonInfoReq) (resp *person_infor.UpdatePersonInfoResp, err error) {
	// Finish your business logic.
	p := &model.PersonInfor{
		UserId: req.UserId,
		Address: model.Address{
			StreetAddress: req.Address.StreetAddress,
			State:         req.Address.State,
			City:          req.Address.City,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
			FirstName:     req.FirstName,
			LastName:      req.LastName,
		},
		CreditCardInfo: model.CreditCardInfo{
			CardNumber:          req.CreditCardInfo.CreditCardNum,
			CVV:                 req.CreditCardInfo.CreitCardCvv,
			CardExpirationYear:  req.CreditCardInfo.CreditCartExpirationYear,
			CardExpirationMonth: req.CreditCardInfo.CreditCartExpirationMonth,
		},
		Email: req.Email,
	}
	err = model.Update(s.ctx, mysql.DB, req.PersonInfoId, p)
	if err != nil {
		return nil, err
	}
	return
}
