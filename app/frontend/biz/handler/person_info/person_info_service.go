package person_info

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/biz/service"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/biz/utils"
	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	person_info "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/person_info"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// PersonInforRegister .
// @router /personinfo/Register [POST]
func PersonInforRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req person_info.PersonInfoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewPersonInforRegisterService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.Redirect(consts.StatusFound, []byte("/personinfo"))
}

// PersonInfo .
// @router /personinfo [GET]
func PersonInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewPersonInfoService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "person-info-list", utils.WarpResponse(ctx, c, resp))
}

// PersonInfoEdit .
// @router /personinfo/edit/:personId [GET]
func PersonInfoEdit(ctx context.Context, c *app.RequestContext) {
	var err error
	var req person_info.PersonInfoEditReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewPersonInfoEditService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.HTML(consts.StatusOK, "person-info", utils.WarpResponse(ctx, c, resp))
}

// PersonInfoDelete .
// @router /personinfo/delete/:personId [GET]
func PersonInfoDelete(ctx context.Context, c *app.RequestContext) {
	var err error
	var req person_info.PersonInfoDeleteReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewPersonInfoDeleteService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.Redirect(consts.StatusFound, []byte("/personinfo"))
}
