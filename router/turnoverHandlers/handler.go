package turnoverHandlers

import (
	"chatgpt-web-new-go/common/auth"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/model"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/turnover"
	"github.com/gin-gonic/gin"
)

func TurnoverListHandler(c *gin.Context) {
	request := &base.Page{}
	if err := c.Bind(request); err != nil {
		base.Fail(c, "参数异常："+err.Error())
		return
	}
	if request.PageSize == 0 {
		request.PageSize = 10
	}
	if request.Page == 0 {
		request.Page = 1
	}

	uFromCtx, found := c.Get(auth.GinCtxKey)
	if !found {
		logs.Error("cannot get auth user key:%v", auth.GinCtxKey)
		base.Fail(c, "cannot get auth user key:"+auth.GinCtxKey)
		return
	}
	u := uFromCtx.(*model.User)

	turnoverList, count, err := turnover.TurnoverList(c, u.ID, request.Page, request.PageSize)
	if err != nil {
		logs.Error("turnover list error: %v", err)
		base.Fail(c, "获取记录异常：%v"+err.Error())
		return
	}

	response := &TurnoverListData{
		Count: count,
		Rows:  turnoverList,
	}
	base.Success(c, response)
}
