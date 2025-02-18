// Code generated by Kitex v0.9.1. DO NOT EDIT.

package orderservice

import (
	"context"
	"errors"
	order "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"ListOrder": kitex.NewMethodInfo(
		listOrderHandler,
		newListOrderArgs,
		newListOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"AddOrder": kitex.NewMethodInfo(
		addOrderHandler,
		newAddOrderArgs,
		newAddOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	orderServiceServiceInfo                = NewServiceInfo()
	orderServiceServiceInfoForClient       = NewServiceInfoForClient()
	orderServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return orderServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return orderServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return orderServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "orderService"
	handlerType := (*order.OrderService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "order",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func listOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.ListOrderReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).ListOrder(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListOrderArgs:
		success, err := handler.(order.OrderService).ListOrder(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListOrderResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListOrderArgs() interface{} {
	return &ListOrderArgs{}
}

func newListOrderResult() interface{} {
	return &ListOrderResult{}
}

type ListOrderArgs struct {
	Req *order.ListOrderReq
}

func (p *ListOrderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.ListOrderReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListOrderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListOrderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListOrderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListOrderArgs) Unmarshal(in []byte) error {
	msg := new(order.ListOrderReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListOrderArgs_Req_DEFAULT *order.ListOrderReq

func (p *ListOrderArgs) GetReq() *order.ListOrderReq {
	if !p.IsSetReq() {
		return ListOrderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListOrderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListOrderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListOrderResult struct {
	Success *order.ListOrderResp
}

var ListOrderResult_Success_DEFAULT *order.ListOrderResp

func (p *ListOrderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.ListOrderResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListOrderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListOrderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListOrderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListOrderResult) Unmarshal(in []byte) error {
	msg := new(order.ListOrderResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListOrderResult) GetSuccess() *order.ListOrderResp {
	if !p.IsSetSuccess() {
		return ListOrderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListOrderResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.ListOrderResp)
}

func (p *ListOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListOrderResult) GetResult() interface{} {
	return p.Success
}

func addOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.AddOrderReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).AddOrder(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *AddOrderArgs:
		success, err := handler.(order.OrderService).AddOrder(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AddOrderResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newAddOrderArgs() interface{} {
	return &AddOrderArgs{}
}

func newAddOrderResult() interface{} {
	return &AddOrderResult{}
}

type AddOrderArgs struct {
	Req *order.AddOrderReq
}

func (p *AddOrderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.AddOrderReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AddOrderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AddOrderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AddOrderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *AddOrderArgs) Unmarshal(in []byte) error {
	msg := new(order.AddOrderReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AddOrderArgs_Req_DEFAULT *order.AddOrderReq

func (p *AddOrderArgs) GetReq() *order.AddOrderReq {
	if !p.IsSetReq() {
		return AddOrderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AddOrderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AddOrderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type AddOrderResult struct {
	Success *order.AddOrderResp
}

var AddOrderResult_Success_DEFAULT *order.AddOrderResp

func (p *AddOrderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.AddOrderResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AddOrderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AddOrderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AddOrderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *AddOrderResult) Unmarshal(in []byte) error {
	msg := new(order.AddOrderResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddOrderResult) GetSuccess() *order.AddOrderResp {
	if !p.IsSetSuccess() {
		return AddOrderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AddOrderResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.AddOrderResp)
}

func (p *AddOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AddOrderResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ListOrder(ctx context.Context, Req *order.ListOrderReq) (r *order.ListOrderResp, err error) {
	var _args ListOrderArgs
	_args.Req = Req
	var _result ListOrderResult
	if err = p.c.Call(ctx, "ListOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddOrder(ctx context.Context, Req *order.AddOrderReq) (r *order.AddOrderResp, err error) {
	var _args AddOrderArgs
	_args.Req = Req
	var _result AddOrderResult
	if err = p.c.Call(ctx, "AddOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
