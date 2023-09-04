package router

import (
	"chatgpt-web-new-go/router/front/authHandlers"
	"chatgpt-web-new-go/router/front/carmiHandlers"
	"chatgpt-web-new-go/router/front/chatHandlers"
	"chatgpt-web-new-go/router/front/configHandlers"
	"chatgpt-web-new-go/router/front/imagesHandlers"
	"chatgpt-web-new-go/router/front/messageHandlers"
	"chatgpt-web-new-go/router/front/payHandlers"
	"chatgpt-web-new-go/router/front/personaHandlers"
	"chatgpt-web-new-go/router/front/pluginHandlers"
	"chatgpt-web-new-go/router/front/productHandlers"
	"chatgpt-web-new-go/router/front/signInHandlers"
	"chatgpt-web-new-go/router/front/turnoverHandlers"
	"chatgpt-web-new-go/router/front/userHandlers"
	"chatgpt-web-new-go/router/middlewares"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	// cors
	r.Use(middlewares.Cors())

	// api root
	apiRoot := r.Group("/api")

	adminGroup(apiRoot.Group("/admin"))

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

	// carmiHandlers
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
	group.GET("", imagesHandlers.ImageList)
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
