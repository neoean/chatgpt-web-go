package userHandlers

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/common/types"
	"chatgpt-web-new-go/model"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/user"
	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	request := &base.Page{}
	if err := c.Bind(request); err != nil {
		logs.Error("user list bind error: %v", err)
		base.Fail(c, "参数异常："+err.Error())
		return
	}

	userList, count, err := user.UserList(c, request.Page, request.PageSize)
	if err != nil {
		logs.Error("user list error: %v", err)
		base.Fail(c, "用户列表查询异常："+err.Error())
		return
	}

	base.Success(c, &base.PageResponse{
		Count: count, Rows: userList,
	})
}

func UserDelete(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" || idStr == "undefined" {
		logs.Error("param empty id: %v", idStr)
		base.Fail(c, "id参数不能为空！")
		return
	}

	id := types.StringToInt64(idStr)

	err := user.UserDelete(c, id)
	if err != nil {
		logs.Error("user delete error: %v", err)
		base.Fail(c, "用户删除失败："+err.Error())
		return
	}

	base.SuccessMsg(c, "删除成功！", nil)
}

func UserUpdate(c *gin.Context) {
	request := &model.User{}

	if err := c.BindJSON(request); err != nil {
		logs.Error("json bind error: %v", err)
		base.Fail(c, "参数异常："+err.Error())
		return
	}

	err := user.UserUpdate(c, request)
	if err != nil {
		logs.Error("user update error: %v", err)
		base.Fail(c, "用户更新失败："+err.Error())
		return
	}

	base.SuccessMsg(c, "更新成功！", nil)
}
