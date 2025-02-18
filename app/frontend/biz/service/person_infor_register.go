package service

import (
	"context"

	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	person_info "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/person_info"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/person_infor"
	"github.com/cloudwego/hertz/pkg/app"
)

type PersonInforRegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPersonInforRegisterService(Context context.Context, RequestContext *app.RequestContext) *PersonInforRegisterService {
	return &PersonInforRegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *PersonInforRegisterService) Run(req *person_info.PersonInfoReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	if req.PersonId > 0 {
		_, err = rpc.PersonInfoClient.UpdatePersonInfo(h.Context, &person_infor.UpdatePersonInfoReq{
			PersonInfoId: req.PersonId,
			UserId:       uint32(frontendUtils.GetUserIdfromCtx(h.Context)),
			Address: &person_infor.Address{
				StreetAddress: req.Street,
				City:          req.City,
				Country:       req.Country,
				ZipCode:       req.Zipcode,
				State:         req.Province,
			},
			Email:     req.Email,
			FirstName: req.Firtname,
			LastName:  req.Lastname,
			CreditCardInfo: &person_infor.CreditCardInfo{
				CreditCardNum:             req.CardNum,
				CreditCartExpirationYear:  req.Year,
				CreditCartExpirationMonth: req.Month,
				CreitCardCvv:              req.Cvv,
			},
		})
		if err != nil {
			return nil, err
		}
		return
	}

	_, err = rpc.PersonInfoClient.RegisterPersonInfo(h.Context, &person_infor.PersonReq{
		UserId: uint32(frontendUtils.GetUserIdfromCtx(h.Context)),
		Address: &person_infor.Address{
			StreetAddress: req.Street,
			City:          req.City,
			Country:       req.Country,
			ZipCode:       req.Zipcode,
			State:         req.Province,
		},
		Email:     req.Email,
		FirstName: req.Firtname,
		LastName:  req.Lastname,
		CreditCardInfo: &person_infor.CreditCardInfo{
			CreditCardNum:             req.CardNum,
			CreditCartExpirationYear:  req.Year,
			CreditCartExpirationMonth: req.Month,
			CreitCardCvv:              req.Cvv,
		},
	})
	if err != nil {
		return nil, err
	}
	return
}
