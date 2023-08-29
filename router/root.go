package router

import (
	"chatgpt-web-new-go/router/authHandlers"
	"chatgpt-web-new-go/router/carmiHandlers"
	"chatgpt-web-new-go/router/chatHandlers"
	"chatgpt-web-new-go/router/configHandlers"
	"chatgpt-web-new-go/router/imagesHandlers"
	"chatgpt-web-new-go/router/messageHandlers"
	"chatgpt-web-new-go/router/middlewares"
	"chatgpt-web-new-go/router/payHandlers"
	"chatgpt-web-new-go/router/personaHandlers"
	"chatgpt-web-new-go/router/pluginHandlers"
	"chatgpt-web-new-go/router/productHandlers"
	"chatgpt-web-new-go/router/signInHandlers"
	"chatgpt-web-new-go/router/turnoverHandlers"
	"chatgpt-web-new-go/router/userHandlers"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	// cors
	r.Use(middlewares.Cors())

	// api root
	apiRoot := r.Group("/api")

	// root group
	authGroup(apiRoot.Group(""))

	// config
	configGroup(apiRoot.Group("/config"))

	// persona
	personaGroup(apiRoot.Group("/persona"))

	// ============== ⬇ need auth ⬇ ==============
	// add jwt
	apiRoot.Use(middlewares.Jwt())

	// user group
	userGroup(apiRoot.Group("/user"))

	// product
	productGroup(apiRoot.Group("/product"))

	// turnover
	turnoverGroup(apiRoot.Group("/turnover"))

	// pay TODO
	payGroup(apiRoot.Group("/pay"))

	// carmi TODO
	carmiGroup(apiRoot.Group("/use_carmi"))

	// signIn
	signInGroup(apiRoot.Group("/signin"))

	// plugin
	pluginGroup(apiRoot.Group("/plugin"))

	// chat group
	chatGroup(apiRoot.Group("/chat"))

	// images group todo
	imagesGroup(apiRoot.Group("/images"))

	// promotion group
	promotionGroup(apiRoot.Group("/promotion"))
}

func pluginGroup(group *gin.RouterGroup) {
	group.GET("", pluginHandlers.PluginHandler)
}

func carmiGroup(group *gin.RouterGroup) {
	group.POST("", carmiHandlers.UseCarmiHandler)
}

func personaGroup(group *gin.RouterGroup) {
	group.GET("", personaHandlers.PersonaHandler)
}

func payGroup(group *gin.RouterGroup) {
	group.POST("/precreate", payHandlers.PayPreCreateHandler)
}

func imagesGroup(group *gin.RouterGroup) {
	group.POST("/generations", imagesHandlers.ImagesGenerationsHandler)
}

func configGroup(group *gin.RouterGroup) {
	group.GET("", configHandlers.ConfigList)
}

func signInGroup(group *gin.RouterGroup) {
	group.POST("", signInHandlers.SignIn)
	group.GET("/list", signInHandlers.SignInList) // 获取用户当月签到记录
}

func turnoverGroup(group *gin.RouterGroup) {
	group.GET("", turnoverHandlers.TurnoverListHandler)
}

func productGroup(group *gin.RouterGroup) {
	group.GET("", productHandlers.ProductListHandler)
}

func authGroup(group *gin.RouterGroup) {
	group.GET("/send_sms", authHandlers.CodeSendHandler)
	group.POST("/login", authHandlers.LoginHandler)
}

func userGroup(group *gin.RouterGroup) {
	group.GET("/info", userHandlers.UserInfoHandler)
	group.GET("/records", userHandlers.UserRecordHandler)
	group.PUT("/password", authHandlers.UserPasswordResetHandler)

	group.GET("/messages", messageHandlers.MessageList)
}

func chatGroup(group *gin.RouterGroup) {
	group.POST("/completions", chatHandlers.ChatCompletionsHandler)
}

func promotionGroup(group *gin.RouterGroup) {
	group.GET("/get_")
}
