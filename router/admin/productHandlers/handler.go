package productHandlers

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/common/types"
	"chatgpt-web-new-go/model"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/product"
	"github.com/gin-gonic/gin"
)

func ProductList(c *gin.Context) {
	request := &base.Page{}

	if err := c.Bind(request); err != nil {
		logs.Error("json bind error: %v", err)
		base.Fail(c, "参数异常："+err.Error())
		return
	}

	productList, count, err := product.ProductList(c, request.Page, request.PageSize)
	if err != nil {
		logs.Error("product lsit error: %v", err)
		base.Fail(c, "商品列表查询异常： "+err.Error())
		return
	}

	base.Success(c, base.PageResponse{
		Count: count,
		Rows:  productList,
	})
}

func ProductAdd(c *gin.Context) {
	request := &model.Product{}
	if err := c.BindJSON(request); err != nil {
		logs.Error("json bind error: %v", err)
		base.Fail(c, "参数异常："+err.Error())
		return
	}

	productAdd, err := product.ProductAdd(c, request)
	if err != nil {
		logs.Error("product add error: %v", err)
		base.Fail(c, "创建商品异常！"+err.Error())
		return
	}

	base.Success(c, productAdd)
}

func ProductPut(c *gin.Context) {
	request := &model.Product{}
	if err := c.BindJSON(request); err != nil {
		logs.Error("json bind error: %v", err)
		base.Fail(c, "参数异常："+err.Error())
		return
	}

	err := product.ProductUpdate(c, request)
	if err != nil {
		logs.Error("product udpate error: %v", err)
		base.Fail(c, "商品更新异常："+err.Error())
		return
	}

	base.SuccessMsg(c, "更新成功", nil)
}

func ProductDelete(c *gin.Context) {
	idStr := c.Param("id")

	if idStr == "" || idStr == "undefined" {
		logs.Error("param id empty: %v", idStr)
		base.Fail(c, "参数id为空："+idStr)
		return
	}

	id := types.StringToInt64(idStr)

	err := product.ProductDelete(c, id)
	if err != nil {
		logs.Error("product delete error: %v", err)
		base.Fail(c, "商品删除异常："+err.Error())
		return
	}

	base.SuccessMsg(c, "删除成功！", nil)
}
