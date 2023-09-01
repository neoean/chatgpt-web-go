package signinHandlers

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/signin"
	"github.com/gin-gonic/gin"
)

func SigninList(c *gin.Context) {
	request := &base.Page{}

	if err := c.Bind(request); err != nil {
		logs.Error("json bind error: %v", err)
		base.Fail(c, "参数异常:"+err.Error())
		return
	}

	signinList, count, err := signin.AdminSigninList(c, request.Page, request.PageSize)
	if err != nil {
		logs.Error("signin list error: %v", err)
		base.Fail(c, "签到列表查询异常："+err.Error())
		return
	}
	base.Success(c, base.PageResponse{
		Count: count,
		Rows:  signinList,
	})
}
