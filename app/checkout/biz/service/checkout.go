package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/biz-demo/gomall/app/checkout/infra/mq"
	"github.com/cloudwego/biz-demo/gomall/app/checkout/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/email"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/person_infor"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	if cartResult == nil || cartResult.Items == nil {
		fmt.Println("cartResult.Items is nil")
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}
	var OrderItems []*order.ProductItems
	var total float32
	var cost float32
	for _, item := range cartResult.Items {
		proRes, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductRequest{Id: item.ProductId})
		if err != nil {
			return nil, err
		}
		if proRes == nil {
			return nil, kerrors.NewGRPCBizStatusError(5004002, "product not found")
		}
		cost = proRes.Product.Price * float32(item.Quantity)
		total += cost
		OrderItems = append(OrderItems, &order.ProductItems{
			Quantity:  item.Quantity,
			ProductId: item.ProductId,
			Cost:      cost,
		})
	}
	PersonInforResp, err := rpc.PersonInfoClient.GetPersonInfo(s.ctx, &person_infor.GetPersonInfoReq{
		PersonInfoId: req.PersonInforId,
	})
	if err != nil {
		return nil, err
	}
	OrdReq := &order.AddOrderReq{
		UserId: req.UserId,
		Email:  PersonInforResp.PersonInfo.Email,
		Address: &checkout.Address{
			StreetAddress: PersonInforResp.PersonInfo.Address.StreetAddress,
			Country:       PersonInforResp.PersonInfo.Address.Country,
			State:         PersonInforResp.PersonInfo.Address.State,
			City:          PersonInforResp.PersonInfo.Address.City,
			ZipCode:       PersonInforResp.PersonInfo.Address.ZipCode,
		},
		Items: OrderItems,
	}
	OrdResp, err := rpc.OrderClient.AddOrder(s.ctx, OrdReq)

	if err != nil {
		return nil, err
	}

	PayReq := &payment.PaymentServiceReq{
		Amount: total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNum:             PersonInforResp.PersonInfo.CreditCardInfo.CreditCardNum,
			CreitCardCvv:              PersonInforResp.PersonInfo.CreditCardInfo.CreitCardCvv,
			CreditCartExpirationYear:  PersonInforResp.PersonInfo.CreditCardInfo.CreditCartExpirationYear,
			CreditCartExpirationMonth: PersonInforResp.PersonInfo.CreditCardInfo.CreditCartExpirationMonth,
		},
		OrderId: OrdResp.OrderId,
		UserId:  req.UserId,
	}
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	PayRes, err := rpc.PaymentClient.Charge(s.ctx, PayReq)
	if err != nil {
		return nil, err
	}
	data, _ := proto.Marshal(&email.EmailReq{
		From:        "admin@shop.com",
		To:          PersonInforResp.PersonInfo.Email,
		Subject:     "Order Confirmation",
		ContentType: "plain/text",
		Content:     "Your order has been placed successfully. Order ID:" + OrdResp.OrderId,
	})
	msg := &nats.Msg{
		Subject: "email",
		Data:    data,
	}
	_ = mq.Nc.PublishMsg(msg)
	return &checkout.CheckoutResp{OrderId: OrdResp.OrderId, TransactionId: PayRes.TransactionId}, nil
}
