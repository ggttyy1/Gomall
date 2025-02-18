package service

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/biz-demo/gomall/app/payment/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/payment/biz/model"
	payment "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.PaymentServiceReq) (resp *payment.PaymentServiceResp, err error) {
	// Finish your business logic.
	fmt.Println("ChargeService")
	// card := creditcard.Card{
	// 	Number: req.CreditCard.CreditCardNum,
	// 	Cvv:    strconv.Itoa(int(req.CreditCard.CreitCardCvv)),
	// 	Month:  strconv.Itoa(int(req.CreditCard.CreditCartExpirationMonth)),
	// 	Year:   strconv.Itoa(int(req.CreditCard.CreditCartExpirationYear)),
	// }
	// err = card.Validate(true)
	// fmt.Println(err.Error())
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	// fmt.Println(err)
	transactionId, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005001, err.Error())
	}
	err = model.CreatePaymentLog(s.ctx, mysql.DB, &model.PaymentLog{
		TransactionId: transactionId.String(),
		OrderId:       req.OrderId,
		UserId:        req.UserId,
		Amount:        req.Amount,
		PayTime:       time.Now(),
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4004002, err.Error())
	}
	return &payment.PaymentServiceResp{
		TransactionId: transactionId.String(),
	}, nil

}
