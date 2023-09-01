package carmiHandlers

import (
	authGlobal "chatgpt-web-new-go/common/auth"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/model"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/carmi"
	"github.com/gin-gonic/gin"
)

func UseCarmiHandler(c *gin.Context) {
	request := &UseCarmiRequest{}

	if err := c.BindJSON(request); err != nil {
		logs.Error("gin bind json error: %v", err)
		base.Fail(c, "参数异常！"+err.Error())
		return
	}

	uFromCtx, found := c.Get(authGlobal.GinCtxKey)
	if !found {
		logs.Error("cannot get auth user key:%v", authGlobal.GinCtxKey)
		base.Fail(c, "cannot get auth user key:"+authGlobal.GinCtxKey)
		return
	}

	u := uFromCtx.(*model.User)

	ip := authGlobal.GetClientIP(c)

	err := carmi.UseCarmi(c, u.ID, request.Carmi, ip)
	if err != nil {
		logs.Error("carmiHandlers use error: %v", err)
		base.Fail(c, "卡密使用异常："+err.Error())
		return
	}

	base.SuccessMsg(c, "充值成功", nil)
}
