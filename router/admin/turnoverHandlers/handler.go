package turnoverHandlers

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/common/types"
	"chatgpt-web-new-go/model"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/turnover"
	"github.com/gin-gonic/gin"
)

func TurnoverList(c *gin.Context) {
	request := &base.Page{}
	if err := c.Bind(request); err != nil {
		logs.Error("request bind error: %v", err)
		base.Fail(c, "参数异常："+err.Error())
		return
	}

	turnoverList, count, err := turnover.AdminTurnoverList(c, request.Page, request.PageSize)
	if err != nil {
		logs.Error("turnover list error: %v", err)
		base.Fail(c, "消费列表查询异常："+err.Error())
		return
	}

	base.Success(c, base.PageResponse{
		Count: count,
		Rows:  turnoverList,
	})
}

func TurnoverDelete(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" || idStr == "undefined" {
		logs.Error("id is empty: %v", idStr)
		base.Fail(c, "参数id不能为空："+idStr)
		return
	}

	id := types.StringToInt64(idStr)
	err := turnover.AdminTurnoverDelete(c, id)
	if err != nil {
		logs.Error("tuanover delete error: %v", err)
		base.Fail(c, "消费记录删除异常！")
		return
	}

	base.SuccessMsg(c, "消费记录删除成功！", nil)
}

func TurnoverUpdate(c *gin.Context) {
	request := &model.Turnover{}

	if err := c.BindJSON(request); err != nil {
		logs.Error("json bind error: %v", err)
		base.Fail(c, "参数异常: "+err.Error())
		return
	}

	turnoverUpdate, err := turnover.AdminTurnoverUpdate(c, request)
	if err != nil {
		logs.Error("tuanover update error: %v", err)
		base.Fail(c, "消费记录更新失败："+err.Error())
		return
	}

	base.Success(c, base.UpdateResponse{turnoverUpdate, true})
}
