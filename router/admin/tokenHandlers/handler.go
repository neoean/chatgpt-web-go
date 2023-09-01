package tokenHandlers

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/common/types"
	"chatgpt-web-new-go/model"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/token"
	"github.com/gin-gonic/gin"
)

func TokenList(c *gin.Context) {
	request := &base.Page{}
	if err := c.Bind(request); err != nil {
		logs.Error("param bind error: %v", err)
		base.Fail(c, "参数异常！"+err.Error())
		return
	}

	tokenList, count, err := token.TokenList(c, request.Page, request.PageSize)
	if err != nil {
		logs.Error("token list error: %v", err)
		base.Fail(c, "token列表查询异常！"+err.Error())
		return
	}

	base.Success(c, base.PageResponse{
		Count: count,
		Rows:  tokenList,
	})
}

func TokenAdd(c *gin.Context) {
	request := &model.Aikey{}
	if err := c.BindJSON(request); err != nil {
		logs.Error("param bind json error: %v", err)
		base.Fail(c, "参数异常："+err.Error())
		return
	}

	tokenAdd, err := token.TokenAdd(c, request)
	if err != nil {
		logs.Error("token add error: %v", err)
		base.Fail(c, "新增Token失败："+err.Error())
		return
	}

	base.Success(c, tokenAdd)
}

func TokenUpdate(c *gin.Context) {
	request := &model.Aikey{}
	if err := c.BindJSON(request); err != nil {
		logs.Error("param bind json error: %v", err)
		base.Fail(c, "参数异常："+err.Error())
		return
	}

	update, err := token.TokenUpdate(c, request)
	if err != nil {
		logs.Error("token update error: %v", err)
		base.Fail(c, "token更新失败："+err.Error())
		return
	}

	base.Success(c, base.UpdateResponse{update, true})
}

func TokenDelete(c *gin.Context) {
	idStr := c.Param("id")

	if idStr == "" || idStr == "undefined" {
		logs.Error("param id is empty: %v", idStr)
		base.Fail(c, "参数id为空！")
		return
	}

	id := types.StringToInt64(idStr)

	err := token.TokenDelete(c, id)
	if err != nil {
		logs.Error("token delete error: %v", err)
		base.Fail(c, "删除异常："+err.Error())
		return
	}

	base.SuccessMsg(c, "删除成功！", nil)
}

func TokenCheck(c *gin.Context) {

}
