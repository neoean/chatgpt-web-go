package amountHandlers

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/common/types"
	"chatgpt-web-new-go/model"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/amount"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	request := &base.Page{}
	if err := c.Bind(request); err != nil {
		logs.Error("json bind error: %v", err)
		base.Fail(c, "参数异常！"+err.Error())
		return
	}

	notifications, count, err := amount.List(c, request.Page, request.PageSize)
	if err != nil {
		logs.Error("amount list error: %v", err)
		base.Fail(c, "查询异常："+err.Error())
		return
	}

	base.Success(c, base.PageResponse{
		Count: count,
		Rows:  notifications,
	})
}

func Update(c *gin.Context) {
	request := &model.AmountDetail{}
	if err := c.BindJSON(request); err != nil {
		logs.Error("json bind error: %v", err)
		base.Fail(c, "参数异常："+err.Error())
		return
	}

	err := amount.Update(c, request)
	if err != nil {
		logs.Error("amount update error: %v", err)
		base.Fail(c, "更新失败："+err.Error())
		return
	}

	base.SuccessMsg(c, "更新成功", nil)
}

func Delete(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" || idStr == "undefined" {
		logs.Error("id is empty: %v", idStr)
		base.Fail(c, "参数id为空！"+idStr)
		return
	}

	id := types.StringToInt64(idStr)
	err := amount.Delete(c, id)
	if err != nil {
		logs.Error("amount delete error: %v", err)
		base.Fail(c, "删除失败！"+err.Error())
		return
	}

	base.SuccessMsg(c, "删除成功！", nil)
}
