package payHandlers

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/common/types"
	"chatgpt-web-new-go/model"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/pay"
	"github.com/gin-gonic/gin"
)

func PayConfigList(c *gin.Context) {
	payments, err := pay.PayConfigList(c)
	if err != nil {
		logs.Error("pay config list error: %v", err)
		base.Fail(c, "支付配置查询异常！"+err.Error())
		return
	}
	base.Success(c, base.PageResponse{
		Count: int64(len(payments)),
		Rows:  payments,
	})
}

func PayConfigAdd(c *gin.Context) {
	request := &model.Payment{}
	if err := c.BindJSON(request); err != nil {
		logs.Error("bind json error: %v", err)
		base.Fail(c, "参数异常："+err.Error())
		return
	}

	err := pay.PayConfigAdd(c, request)
	if err != nil {
		logs.Error("pay config create error: %v", err)
		base.Fail(c, "支付配置新增异常："+err.Error())
		return
	}

	base.SuccessMsg(c, "支付配置新增成功！", nil)
}

func PayConfigUpdate(c *gin.Context) {
	request := &model.Payment{}
	if err := c.BindJSON(request); err != nil {
		logs.Error("json bind error:  %v", err)
		base.Fail(c, "参数异常！"+err.Error())
		return
	}

	err := pay.PayConfigUpdate(c, request)
	if err != nil {
		logs.Error("pay config update error: %v", err)
		base.Fail(c, "支付配置更新异常："+err.Error())
		return
	}

	base.SuccessMsg(c, "更新成功！", nil)
}

func PayConfigDelete(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" || idStr == "undefined" {
		logs.Error("id is empty: %v", idStr)
		base.Fail(c, "参数id为空！"+idStr)
		return
	}

	id := types.StringToInt64(idStr)
	err := pay.PayConfigDelete(c, id)
	if err != nil {
		logs.Error("pay config delete error: %v", err)
		base.Fail(c, "支付配置删除失败！"+err.Error())
		return
	}

	base.SuccessMsg(c, "删除成功！", nil)
}
