package service

import (
	"context"
	"testing"

	payment "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/payment"
)

func TestCharge_Run(t *testing.T) {
	ctx := context.Background()
	s := NewChargeService(ctx)
	// init req and assert value
	req := &payment.PaymentServiceReq{
		Amount: 123,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNum:             "1234567890123456",
			CreitCardCvv:              123,
			CreditCartExpirationYear:  123,
			CreditCartExpirationMonth: 123,
		},
		OrderId: "1",
		UserId:  1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
