package orderHandlers

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/order"
	"github.com/gin-gonic/gin"
)

func OrderList(c *gin.Context) {
	request := &base.Page{}
	if err := c.Bind(request); err != nil {
		logs.Error("json bind error: %v", err)
		base.Fail(c, "参数异常: "+err.Error())
		return
	}

	orderList, count, err := order.OrderList(c, request.Page, request.PageSize)
	if err != nil {
		logs.Error("order list error: %v", err)
		base.Fail(c, "订单查询失败："+err.Error())
		return
	}

	base.Success(c, base.PageResponse{
		Count: count,
		Rows:  orderList,
	})
}
