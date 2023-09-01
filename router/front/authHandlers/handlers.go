package authHandlers

import (
	authGlobal "chatgpt-web-new-go/common/auth"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/common/types"
	"chatgpt-web-new-go/model"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/auth"
	"chatgpt-web-new-go/service/signin"
	"chatgpt-web-new-go/service/sns"
	"github.com/gin-gonic/gin"
)

func CodeSendHandler(c *gin.Context) {
	source := c.Query("source")
	if source == "" {
		base.Fail(c, "账号不能为空!")
		return
	}

	err := sns.SendCode(source)
	if err != nil {
		base.Fail(c, "发送失败！"+err.Error())
		return
	}

	base.SuccessMsg(c, "发送成功，请注意查收！", nil)
}

func LoginHandler(c *gin.Context) {
	request := &CodeSendRequest{}
	if err := c.BindJSON(request); err != nil {
		base.Fail(c, "账号参数异常："+err.Error())
		return
	}
	if request.Account == "" || (request.Code == "" && request.Password == "") {
		base.Fail(c, "登陆参数不能为空!")
		return
	}

	u, token, err := auth.Login(c, request.Account, request.Password, request.Code, request.InviteCode)
	if err != nil {
		base.Fail(c, "登陆失败！"+err.Error())
		return
	}

	isSignIn := signin.IsSignInToday(u.ID)
	response := &LoginResponse{
		UserInfo: UserInfo{
			User:     u,
			IsSignIn: types.Booltonumber(isSignIn),
		},
		Token: token,
	}
	base.SuccessMsg(c, "登陆成功", response)
}

func SessionHandler(c *gin.Context) {
	response := &SessionResponse{
		Auth:  true,
		Model: ModelOfficial,
	}
	base.Success(c, response)
}

func UserPasswordResetHandler(c *gin.Context) {
	uFromCtx, found := c.Get(authGlobal.GinCtxKey)
	if !found {
		logs.Error("cannot get auth user key:%v", authGlobal.GinCtxKey)
		base.Fail(c, "cannot get auth user key:"+authGlobal.GinCtxKey)
		return
	}

	u := uFromCtx.(*model.User)

	request := &CodeSendRequest{}
	if err := c.BindJSON(request); err != nil {
		base.Fail(c, "账号参数异常："+err.Error())
		return
	}
	if request.Account == "" || (request.Code == "" && request.Password == "") {
		base.Fail(c, "参数不能为空!")
		return
	}

	if request.Account != u.Account {
		base.Fail(c, "非法修改操作!")
		return
	}

	err := auth.UserPasswordReset(c, request.Account, request.Code, request.Password)
	if err != nil {
		base.Fail(c, "密码修改失败!"+err.Error())
		return
	}

	base.SuccessMsg(c, "密码修改成功！", nil)
}
