package ai

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/biz/service"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/biz/utils"
	ai "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/ai"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetAiInfo .
// @router /ai/search [GET]
func GetAiInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ai.GetAiInfoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetAiInfoService(ctx, c).Run(&req)
	fmt.Println(resp["url"])
	urlValue, ok := resp["url"].(string)
	fmt.Println(urlValue)
	if !ok {
		// 处理类型断言失败的情况
		log.Fatal("resp[\"url\"] is not of type []byte")
	}
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.Redirect(consts.StatusOK, []byte(urlValue[1:len(urlValue)-1]))
}
