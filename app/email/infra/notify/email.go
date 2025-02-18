package notify

import (
	"time"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/email"
	"github.com/kr/pretty"
)

type NoodEmail struct {
}

func (e *NoodEmail) SendEmail(req *email.EmailReq) error {
	time.Sleep(10 * time.Second)
	pretty.Printf("%v\n", req)
	return nil
}

func NewNoodEmail() NoodEmail {
	return NoodEmail{}
}
