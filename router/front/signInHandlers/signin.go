package signInHandlers

import (
	"chatgpt-web-new-go/common/auth"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/model"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/signin"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	uFromCtx, found := c.Get(auth.GinCtxKey)
	if !found {
		logs.Error("cannot get auth user key:%v", auth.GinCtxKey)
		base.Fail(c, "请先登陆")
		return
	}

	u := uFromCtx.(*model.User)
	ip := auth.GetClientIP(c)

	err := signin.SignIn(c, u.ID, ip)
	if err != nil {
		logs.Error("signin error: %v", err)
		base.Fail(c, "签到异常: "+err.Error())
		return
	}

	base.SuccessMsg(c, fmt.Sprintf("签到成功 +%v积分", signin.SigninCoin), nil)
}

func SignInList(c *gin.Context) {
	uFromCtx, found := c.Get(auth.GinCtxKey)
	if !found {
		logs.Error("cannot get auth user key:%v", auth.GinCtxKey)
		base.Fail(c, "请先登陆")
		return
	}

	u := uFromCtx.(*model.User)

	signinList, err := signin.GetSignListMonth(u.ID)
	if err != nil {
		logs.Error("signin.GetSignListMonth error: %v", err)
		base.Fail(c, "获取签到记录异常！"+err.Error())
		return
	}

	base.Success(c, signinList)
}
