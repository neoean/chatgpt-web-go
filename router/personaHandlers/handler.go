package personaHandlers

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/persona"
	"github.com/gin-gonic/gin"
)

func PersonaHandler(c *gin.Context) {

	personaList, err := persona.PersonaList(c)
	if err != nil {
		logs.Error("persona list error: %v", err)
		base.Fail(c, "查询异常："+err.Error())
		return
	}

	base.Success(c, personaList)
}
