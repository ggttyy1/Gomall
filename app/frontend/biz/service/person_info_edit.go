package service

import (
	"context"
	"strconv"

	person_info "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/person_info"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/person_infor"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type PersonInfoEditService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPersonInfoEditService(Context context.Context, RequestContext *app.RequestContext) *PersonInfoEditService {
	return &PersonInfoEditService{RequestContext: RequestContext, Context: Context}
}

func (h *PersonInfoEditService) Run(req *person_info.PersonInfoEditReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	PersonInfoResp, err := rpc.PersonInfoClient.GetPersonInfo(h.Context, &person_infor.GetPersonInfoReq{
		PersonInfoId: req.PersonId,
	})
	if err != nil {
		return nil, err
	}
	personItems := map[string]string{
		"StreetAddress":   PersonInfoResp.PersonInfo.Address.StreetAddress,
		"City":            PersonInfoResp.PersonInfo.Address.City,
		"State":           PersonInfoResp.PersonInfo.Address.State,
		"ZipCode":         PersonInfoResp.PersonInfo.Address.ZipCode,
		"Country":         PersonInfoResp.PersonInfo.Address.Country,
		"CreditCardNum":   PersonInfoResp.PersonInfo.CreditCardInfo.CreditCardNum,
		"CreditCardYeat":  strconv.Itoa(int(PersonInfoResp.PersonInfo.CreditCardInfo.CreditCartExpirationYear)),
		"CreditCardMonth": strconv.Itoa(int(PersonInfoResp.PersonInfo.CreditCardInfo.CreditCartExpirationMonth)),
		"Cvv":             strconv.Itoa(int(PersonInfoResp.PersonInfo.CreditCardInfo.CreitCardCvv)),
		"Firstname":       PersonInfoResp.PersonInfo.FirstName,
		"Lastname":        PersonInfoResp.PersonInfo.LastName,
		"Email":           PersonInfoResp.PersonInfo.Email,
		"PersonId":        strconv.Itoa(int(PersonInfoResp.PersonInfo.PersonInfoId)),
	}

	return utils.H{
		"Title": "PersonInfoEdit",
		"items": personItems,
	}, nil
}
