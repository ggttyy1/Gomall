package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/person_infor/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/person_infor/biz/model"
	person_infor "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/person_infor"
)

type GetPersonInfoService struct {
	ctx context.Context
} // NewGetPersonInfoService new GetPersonInfoService
func NewGetPersonInfoService(ctx context.Context) *GetPersonInfoService {
	return &GetPersonInfoService{ctx: ctx}
}

// Run create note info
func (s *GetPersonInfoService) Run(req *person_infor.GetPersonInfoReq) (resp *person_infor.GetPersonInfoResp, err error) {
	// Finish your business logic.
	PersonInfo, err := model.SearchPersonInforByPersonId(s.ctx, mysql.DB, req.PersonInfoId)
	if err != nil {
		return nil, err
	}
	resp = &person_infor.GetPersonInfoResp{
		PersonInfo: &person_infor.PersonInfo{
			PersonInfoId: uint32(PersonInfo.ID),
			FirstName:    PersonInfo.Address.FirstName,
			LastName:     PersonInfo.Address.LastName,
			Email:        PersonInfo.Email,
			Address: &person_infor.Address{
				Country:       PersonInfo.Address.Country,
				City:          PersonInfo.Address.City,
				StreetAddress: PersonInfo.Address.StreetAddress,
				State:         PersonInfo.Address.State,
				ZipCode:       PersonInfo.Address.ZipCode,
			},
			CreditCardInfo: &person_infor.CreditCardInfo{
				CreditCardNum:             PersonInfo.CreditCardInfo.CardNumber,
				CreitCardCvv:              PersonInfo.CreditCardInfo.CVV,
				CreditCartExpirationYear:  PersonInfo.CreditCardInfo.CardExpirationYear,
				CreditCartExpirationMonth: PersonInfo.CreditCardInfo.CardExpirationMonth,
			},
		},
	}

	return
}
