package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/person_infor/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/person_infor/biz/model"
	person_infor "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/person_infor"
)

type RegisterPersonInfoService struct {
	ctx context.Context
} // NewRegisterPersonInfoService new RegisterPersonInfoService
func NewRegisterPersonInfoService(ctx context.Context) *RegisterPersonInfoService {
	return &RegisterPersonInfoService{ctx: ctx}
}

// Run create note info
func (s *RegisterPersonInfoService) Run(req *person_infor.PersonReq) (resp *person_infor.PersonResp, err error) {
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
	err = model.Create(s.ctx, mysql.DB, p)
	if err != nil {
		return nil, err
	}
	return
}
