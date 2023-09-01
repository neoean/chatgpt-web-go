package configHandlers

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/config"
	"github.com/gin-gonic/gin"
)

func ConfigList(c *gin.Context) {
	configList, err := config.AdminConfigList(c)
	if err != nil {
		logs.Error("config list error: %v", err)
		base.Fail(c, "配置列表查询异常："+err.Error())
		return
	}

	base.Success(c, configList)
}

func ConfigUpdate(c *gin.Context) {
	request := map[string]string{}
	if err := c.BindJSON(&request); err != nil {
		logs.Error("param bind json error: %v", err)
		base.Fail(c, "参数异常！"+err.Error())
		return
	}

	err := config.ConfigUpdate(c, request)
	if err != nil {
		logs.Error("config update error: %v", err)
		base.Fail(c, "配置更新异常！"+err.Error())
		return
	}

	base.SuccessMsg(c, "更新成功！", nil)
}
