package messageHandlers

import (
	authGlobal "chatgpt-web-new-go/common/auth"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/model"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/message"
	"github.com/gin-gonic/gin"
)

func MessageList(c *gin.Context) {
	uFromCtx, found := c.Get(authGlobal.GinCtxKey)
	if !found {
		logs.Error("cannot get auth user key:%v", authGlobal.GinCtxKey)
		base.Fail(c, "cannot get auth user key:"+authGlobal.GinCtxKey)
		return
	}

	u := uFromCtx.(*model.User)

	messageList, err := message.MessageList(c, u.ID)
	if err != nil {
		logs.Error("message list error:%v", err)
		base.Fail(c, "查询异常:"+err.Error())
		return
	}

	if len(messageList) == 0 {
		messageList = make([]*message.Message, 0)
	}

	base.Success(c, messageList)
}
