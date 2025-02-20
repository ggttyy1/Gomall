package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/biz/utils/tools"
	ai "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/ai"
	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetAiInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetAiInfoService(Context context.Context, RequestContext *app.RequestContext) *GetAiInfoService {
	return &GetAiInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *GetAiInfoService) Run(req *ai.GetAiInfoReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	// ctx := h.Context
	// rp, err := tools.Agent.Invoke(ctx, []*schema.Message{
	// 	{
	// 		Role:    schema.User,
	// 		Content: req.Question,
	// 	},
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println("complete")
	// // 输出结果
	// for idx, msg := range rp {
	// 	tools.Infof("\n")
	// 	tools.Infof("message %d: %s: %s", idx, msg.Role, msg.Content)
	// }
	rp, err := tools.Agent.Invoke(h.Context, []*schema.Message{
		{
			Role:    schema.User,
			Content: req.Question,
		},
	})
	resp = utils.H{}
	for _, msg := range rp {
		resp["url"] = msg.Content

	}
	return resp, err
}
