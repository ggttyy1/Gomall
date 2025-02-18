package person_infor

import (
	"context"
	person_infor "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/person_infor"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/person_infor/personinfor"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() personinfor.Client
	Service() string
	RegisterPersonInfo(ctx context.Context, Req *person_infor.PersonReq, callOptions ...callopt.Option) (r *person_infor.PersonResp, err error)
	GetPersonInfoList(ctx context.Context, Req *person_infor.GetPersonInfoListReq, callOptions ...callopt.Option) (r *person_infor.GetPersonInfoListResp, err error)
	GetPersonInfo(ctx context.Context, Req *person_infor.GetPersonInfoReq, callOptions ...callopt.Option) (r *person_infor.GetPersonInfoResp, err error)
	UpdatePersonInfo(ctx context.Context, Req *person_infor.UpdatePersonInfoReq, callOptions ...callopt.Option) (r *person_infor.UpdatePersonInfoResp, err error)
	DeletePersonInfo(ctx context.Context, Req *person_infor.DeletePersonInfoReq, callOptions ...callopt.Option) (r *person_infor.DeletePersonInfoResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := personinfor.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient personinfor.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() personinfor.Client {
	return c.kitexClient
}

func (c *clientImpl) RegisterPersonInfo(ctx context.Context, Req *person_infor.PersonReq, callOptions ...callopt.Option) (r *person_infor.PersonResp, err error) {
	return c.kitexClient.RegisterPersonInfo(ctx, Req, callOptions...)
}

func (c *clientImpl) GetPersonInfoList(ctx context.Context, Req *person_infor.GetPersonInfoListReq, callOptions ...callopt.Option) (r *person_infor.GetPersonInfoListResp, err error) {
	return c.kitexClient.GetPersonInfoList(ctx, Req, callOptions...)
}

func (c *clientImpl) GetPersonInfo(ctx context.Context, Req *person_infor.GetPersonInfoReq, callOptions ...callopt.Option) (r *person_infor.GetPersonInfoResp, err error) {
	return c.kitexClient.GetPersonInfo(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdatePersonInfo(ctx context.Context, Req *person_infor.UpdatePersonInfoReq, callOptions ...callopt.Option) (r *person_infor.UpdatePersonInfoResp, err error) {
	return c.kitexClient.UpdatePersonInfo(ctx, Req, callOptions...)
}

func (c *clientImpl) DeletePersonInfo(ctx context.Context, Req *person_infor.DeletePersonInfoReq, callOptions ...callopt.Option) (r *person_infor.DeletePersonInfoResp, err error) {
	return c.kitexClient.DeletePersonInfo(ctx, Req, callOptions...)
}
