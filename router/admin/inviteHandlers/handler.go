package inviteHandlers

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/common/types"
	"chatgpt-web-new-go/model"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/invite"
	"github.com/gin-gonic/gin"
)

func InviteRecords(c *gin.Context) {
	request := &base.Page{}
	if err := c.Bind(request); err != nil {
		logs.Error("bind json error: %v", err)
		base.Fail(c, "参数异常："+err.Error())
		return
	}

	inviteRecords, count, err := invite.InviteRecords(c, request.Page, request.PageSize)
	if err != nil {
		logs.Error("invite records error: %v", err)
		base.Fail(c, "邀请记录查询异常！"+err.Error())
		return
	}

	base.Success(c, base.PageResponse{
		Count: count,
		Rows:  inviteRecords,
	})
}

func InviteUpdate(c *gin.Context) {
	request := &model.InviteRecord{}

	if err := c.BindJSON(request); err != nil {
		logs.Error("json bind error: %v", err)
		base.Fail(c, "参数异常："+err.Error())
		return
	}

	err := invite.InviteUpdate(c, request)
	if err != nil {
		logs.Error("invite update error: %v", err)
		base.Fail(c, "邀请记录更新异常："+err.Error())
		return
	}

	base.SuccessMsg(c, "更新成功！", nil)
}

func InviteDelete(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" || idStr == "undefined" {
		logs.Error("param id not found: %v", idStr)
		base.Fail(c, "缺少参数id")
		return
	}

	id := types.StringToInt64(idStr)
	err := invite.InviteDelete(c, id)
	if err != nil {
		logs.Error("delete invite error: %v", err)
		base.Fail(c, "删除失败："+err.Error())
		return
	}

	base.SuccessMsg(c, "删除成功", nil)
}

func InvitePass(c *gin.Context) {
	err := invite.InvitePass(c)
	if err != nil {
		logs.Error("invite pass error: %v", err)
		base.Fail(c, "邀请记录通过异常！"+err.Error())
		return
	}

	base.SuccessMsg(c, "全部通过成功！", nil)
}
