package imagesHandlers

import (
	authGlobal "chatgpt-web-new-go/common/auth"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/model"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/draw"
	"github.com/gin-gonic/gin"
)

func ImageList(c *gin.Context) {
	request := &ImagesRequest{}

	if err := c.Bind(request); err != nil {
		logs.Error("request bind error: %v", err)
		base.Fail(c, "参数异常:"+err.Error())
		return
	}

	uFromCtx, found := c.Get(authGlobal.GinCtxKey)
	if !found {
		logs.Error("cannot get auth user key:%v", authGlobal.GinCtxKey)
		base.Fail(c, "cannot get auth user key:"+authGlobal.GinCtxKey)
		return
	}

	u := uFromCtx.(*model.User)

	drawRecords, count, err := draw.DrawList(c, u.ID, request.Type, request.Page, request.PageSize)
	if err != nil {
		logs.Error("draw list error: %v", err)
		base.Fail(c, "查询异常："+err.Error())
		return
	}

	base.Success(c, base.PageResponse{
		Count: count,
		Rows:  drawRecords,
	})
}

func ImagesGenerationsHandler(c *gin.Context) {

}
