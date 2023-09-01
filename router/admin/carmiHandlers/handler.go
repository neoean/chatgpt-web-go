package carmiHandlers

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/common/types"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/carmi"
	"github.com/gin-gonic/gin"
)

func CarmiList(c *gin.Context) {
	request := &base.Page{}

	if err := c.BindJSON(request); err != nil {
		logs.Error("request bind json error: %v", err)
		base.Fail(c, "参数异常:"+err.Error())
		return
	}

	carmiList, count, err := carmi.CarmiList(c, request.Page, request.PageSize)
	if err != nil {
		logs.Error("carmi list error: %v", err)
		base.Fail(c, "查询异常："+err.Error())
		return
	}

	base.Success(c, &base.PageResponse{
		Count: count, Rows: carmiList,
	})
}

func CarmiDel(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" || idStr == "undefined" {
		logs.Error("param id not found: %v", idStr)
		base.Fail(c, "缺少参数id")
		return
	}

	id := types.StringToInt64(idStr)
	err := carmi.CarmiDel(c, id)
	if err != nil {
		logs.Error("delete carmi error: %v", err)
		base.Fail(c, "删除失败："+err.Error())
		return
	}

	base.SuccessMsg(c, "删除成功", nil)
}

func CarmiGen(c *gin.Context) {
	request := &carmi.CarmiGenRequest{}

	if err := c.BindJSON(request); err != nil {
		logs.Error("request bind error: %v", err)
		base.Fail(c, "参数异常："+err.Error())
		return
	}

	result, err := carmi.CarmiGen(c, request)
	if err != nil {
		logs.Error("carmi gen service error: %v", err)
		base.Fail(c, "卡密生成异常："+err.Error())
		return
	}

	base.Success(c, result)
}

func CarmiCheck(c *gin.Context) {

}
