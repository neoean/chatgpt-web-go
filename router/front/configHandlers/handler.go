package configHandlers

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/config"
	"github.com/gin-gonic/gin"
)

func ConfigList(c *gin.Context) {
	configList, err := config.ConfigList(c)
	if err != nil {
		logs.Error("config list error: %v", err)
		base.Fail(c, "获取配置信息异常："+err.Error())
		return
	}

	base.Success(c, configList)
}
