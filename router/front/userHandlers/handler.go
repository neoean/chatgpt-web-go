package userHandlers

import (
	"chatgpt-web-new-go/common/auth"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/model"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/user"
	"github.com/gin-gonic/gin"
)

func UserInfoHandler(c *gin.Context) {
	uFromCtx, found := c.Get(auth.GinCtxKey)
	if !found {
		logs.Error("cannot get auth user key:%v", auth.GinCtxKey)
		base.Fail(c, "cannot get auth user key:"+auth.GinCtxKey)
		return
	}

	u := uFromCtx.(*model.User)

	userModel, err := user.GetUserInfo(c, u.Account)
	if err != nil {
		logs.Error("user GetUserInfo bizError: %v", err)
		base.Fail(c, "user GetUserInfo bizError")
		return
	}
	base.Success(c, userModel)
}

func UserRecordHandler(c *gin.Context) {

}
