package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/biz-demo/gomall/app/person_infor/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/person_infor/biz/model"
	person_infor "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/person_infor"
)

type GetPersonInfoListService struct {
	ctx context.Context
} // NewGetPersonInfoListService new GetPersonInfoListService
func NewGetPersonInfoListService(ctx context.Context) *GetPersonInfoListService {
	return &GetPersonInfoListService{ctx: ctx}
}

// Run create note info
func (s *GetPersonInfoListService) Run(req *person_infor.GetPersonInfoListReq) (resp *person_infor.GetPersonInfoListResp, err error) {
	// Finish your business logic.
	PersonIforList, err := model.SearchPersonInforByUserId(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, err
	}
	var PersonResp []*person_infor.PersonInfo
	for _, PersonInfo := range PersonIforList {
		PersonItem := &person_infor.PersonInfo{
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
		}
		PersonResp = append(PersonResp, PersonItem)
	}
	fmt.Println(PersonResp)
	return &person_infor.GetPersonInfoListResp{PersonInfo: PersonResp}, nil
}
