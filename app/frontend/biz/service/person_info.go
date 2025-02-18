package service

import (
	"context"
	"strconv"

	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/person_infor"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type PersonInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPersonInfoService(Context context.Context, RequestContext *app.RequestContext) *PersonInfoService {
	return &PersonInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *PersonInfoService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userid := frontendUtils.GetUserIdfromCtx(h.Context)
	PersonInfoResp, err := rpc.PersonInfoClient.GetPersonInfoList(h.Context, &person_infor.GetPersonInfoListReq{
		UserId: uint32(userid),
	})
	if err != nil {
		return nil, err
	}
	var personItems []map[string]string
	for _, v := range PersonInfoResp.PersonInfo {
		personItems = append(personItems, map[string]string{
			"StreetAddress":   v.Address.StreetAddress,
			"City":            v.Address.City,
			"State":           v.Address.State,
			"ZipCode":         v.Address.ZipCode,
			"Country":         v.Address.Country,
			"CreditCardNum":   v.CreditCardInfo.CreditCardNum,
			"CreditCardYeat":  strconv.Itoa(int(v.CreditCardInfo.CreditCartExpirationYear)),
			"CreditCardMonth": strconv.Itoa(int(v.CreditCardInfo.CreditCartExpirationMonth)),
			"Cvv":             strconv.Itoa(int(v.CreditCardInfo.CreitCardCvv)),
			"Firstname":       v.FirstName,
			"Lastname":        v.LastName,
			"Email":           v.Email,
			"PersonInforId":   strconv.Itoa(int(v.PersonInfoId)),
		})
	}
	return utils.H{
		"Title": "PersonInfo",
		"items": personItems,
	}, nil
}
