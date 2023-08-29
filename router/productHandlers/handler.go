package productHandlers

import (
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/pay"
	"chatgpt-web-new-go/service/product"
	"github.com/gin-gonic/gin"
)

// ProductListHandler product list
func ProductListHandler(c *gin.Context) {
	request := &base.Page{}
	if err := c.Bind(request); err != nil {
		base.Fail(c, "参数异常："+err.Error())
		return
	}

	if request.PageSize == 0 {
		request.PageSize = 10
	}

	productList, _, err := product.ProductList(c, request.Page, request.PageSize)
	if err != nil {
		base.Fail(c, "商品查询异常："+err.Error())
		return
	}

	payTypeList, err := pay.PayTypeList(c)
	if err != nil {
		base.Fail(c, "支付信息查询异常："+err.Error())
		return
	}

	response := &ProductListData{
		Products: productList,
		PayTypes: payTypeList,
	}

	base.Success(c, response)
}
