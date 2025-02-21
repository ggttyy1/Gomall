package service

import (
	"context"
	"strconv"

	checkout "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/checkout"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	rpccart "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/person_infor"
	rpcproduct "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutProductsService(Context context.Context, RequestContext *app.RequestContext) *CheckoutProductsService {
	return &CheckoutProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutProductsService) Run(req *checkout.CheckoutProductReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	var items []map[string]string
	userid := frontendUtils.GetUserIdfromCtx(h.Context)

	carts, err := rpc.CartClient.GetCartByProduct(h.Context, &rpccart.GetCartByProductReq{UserId: uint32(userid), ProductIds: req.SelectedProduct})
	if err != nil {
		return nil, err
	}
	var total float64
	for _, cart := range carts.Items {
		product, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductRequest{Id: cart.ProductId})
		if err != nil {
			return nil, err
		}
		if product.Product == nil {
			continue
		}
		p := product.Product
		items = append(items, map[string]string{
			"Name":      p.Name,
			"Price":     strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture":   p.Picture,
			"Qty":       strconv.Itoa(int(cart.Quantity)),
			"ProductId": strconv.Itoa(int(cart.ProductId)),
		})
		total += float64(p.Price) * float64(cart.Quantity)
	}
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
		"title":       "checkout",
		"items":       items,
		"personItems": personItems,
		"total":       strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
