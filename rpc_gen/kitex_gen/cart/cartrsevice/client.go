// Code generated by Kitex v0.9.1. DO NOT EDIT.

package cartrsevice

import (
	"context"
	cart "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AddItem(ctx context.Context, Req *cart.AddItemReq, callOptions ...callopt.Option) (r *cart.AddItemResp, err error)
	GetCart(ctx context.Context, Req *cart.GetCartReq, callOptions ...callopt.Option) (r *cart.GetCartResp, err error)
	EmptyCart(ctx context.Context, Req *cart.EmptyCartReq, callOptions ...callopt.Option) (r *cart.EmptyCartResp, err error)
	GetCartByProduct(ctx context.Context, Req *cart.GetCartByProductReq, callOptions ...callopt.Option) (r *cart.GetCartByProductResp, err error)
	RemoveCartByProduct(ctx context.Context, Req *cart.RemoveCartByProductReq, callOptions ...callopt.Option) (r *cart.RemoveCartByProductResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kCartrSeviceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCartrSeviceClient struct {
	*kClient
}

func (p *kCartrSeviceClient) AddItem(ctx context.Context, Req *cart.AddItemReq, callOptions ...callopt.Option) (r *cart.AddItemResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddItem(ctx, Req)
}

func (p *kCartrSeviceClient) GetCart(ctx context.Context, Req *cart.GetCartReq, callOptions ...callopt.Option) (r *cart.GetCartResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCart(ctx, Req)
}

func (p *kCartrSeviceClient) EmptyCart(ctx context.Context, Req *cart.EmptyCartReq, callOptions ...callopt.Option) (r *cart.EmptyCartResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EmptyCart(ctx, Req)
}

func (p *kCartrSeviceClient) GetCartByProduct(ctx context.Context, Req *cart.GetCartByProductReq, callOptions ...callopt.Option) (r *cart.GetCartByProductResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCartByProduct(ctx, Req)
}

func (p *kCartrSeviceClient) RemoveCartByProduct(ctx context.Context, Req *cart.RemoveCartByProductReq, callOptions ...callopt.Option) (r *cart.RemoveCartByProductResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RemoveCartByProduct(ctx, Req)
}
