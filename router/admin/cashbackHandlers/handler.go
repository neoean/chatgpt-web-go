package cashbackHandlers

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/common/types"
	"chatgpt-web-new-go/model"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/cashback"
	"github.com/gin-gonic/gin"
)

func CashbackList(c *gin.Context) {
	request := &base.Page{}
	if err := c.Bind(request); err != nil {
		logs.Error("json bind error: %v", err)
		base.Fail(c, "参数异常："+err.Error())
		return
	}

	cashbackList, count, err := cashback.CashbackList(c, request.Page, request.PageSize)
	if err != nil {
		logs.Error("cashback list error: %v", err)
		base.Fail(c, "查询异常！"+err.Error())
		return
	}

	base.Success(c, base.PageResponse{
		Count: count,
		Rows:  cashbackList,
	})
}

func CashbackUpdate(c *gin.Context) {
	request := &model.Cashback{}
	if err := c.BindJSON(request); err != nil {
		logs.Error("json bind error: %v", err)
		base.Fail(c, "参数异常！"+err.Error())
		return
	}

	err := cashback.CashbackUpdate(c, request)
	if err != nil {
		logs.Error("cashback update error: %v", err)
		base.Fail(c, "更新异常："+err.Error())
		return
	}

	base.SuccessMsg(c, "更新成功！", nil)
}

func CashbackDelete(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" || idStr == "undefined" {
		logs.Error("param id not found: %v", idStr)
		base.Fail(c, "缺少参数id")
		return
	}

	id := types.StringToInt64(idStr)
	err := cashback.CashbackDelete(c, id)
	if err != nil {
		logs.Error("delete cashback error: %v", err)
		base.Fail(c, "删除失败："+err.Error())
		return
	}

	base.SuccessMsg(c, "删除成功", nil)
}

func CashbackPass(c *gin.Context) {
	request := &model.Cashback{}
	if err := c.BindJSON(request); err != nil {
		logs.Error("json bind error: %v", err)
		base.Fail(c, "参数异常！"+err.Error())
		return
	}

	err := cashback.CashbackPass(c, request.ID)
	if err != nil {
		logs.Error("cashback pass error: %v", err)
		base.Fail(c, "更新异常："+err.Error())
		return
	}

	base.SuccessMsg(c, "更新成功！", nil)
}
