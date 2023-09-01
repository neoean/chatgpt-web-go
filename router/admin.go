package router

import (
	"chatgpt-web-new-go/router/admin/carmiHandlers"
	"chatgpt-web-new-go/router/admin/cashbackHandlers"
	"chatgpt-web-new-go/router/admin/configHandlers"
	"chatgpt-web-new-go/router/admin/inviteHandlers"
	"chatgpt-web-new-go/router/admin/messageHandlers"
	"chatgpt-web-new-go/router/admin/notificationHandlers"
	"chatgpt-web-new-go/router/admin/orderHandlers"
	"chatgpt-web-new-go/router/admin/payHandlers"
	"chatgpt-web-new-go/router/admin/productHandlers"
	"chatgpt-web-new-go/router/admin/signinHandlers"
	"chatgpt-web-new-go/router/admin/tokenHandlers"
	"chatgpt-web-new-go/router/admin/turnoverHandlers"
	"chatgpt-web-new-go/router/admin/userHandlers"
	"github.com/gin-gonic/gin"
)

func adminGroup(group *gin.RouterGroup) {

	// carmi Handlers
	adminCarmiGroup(group.Group("/carmi"))

	// user handlers
	adminUserGroup(group.Group("/user"))

	// product
	adminProductGroup(group.Group("/products"))

	// message
	adminMessageGroup(group.Group("/messages"))

	// signin
	adminSigninGroup(group.Group("/signin"))

	// turnover
	adminTurnoverGroup(group.Group("/turnover"))

	// token
	adminTokenGroup(group.Group("/aikey"))

	// config
	adminConfigGroup(group.Group("/config"))

	// pay
	adminPayGroup(group.Group("/payment"))

	// orders
	adminOrderGroup(group.Group("/orders"))

	// notice
	adminNotificationGroup(group.Group("/notification"))

	// invite_record
	adminInviteGroup(group.Group("/invite_record"))

	// cashback
	adminCashbackGroup(group.Group("/cashback"))
}

func adminCashbackGroup(group *gin.RouterGroup) {
	group.GET("", cashbackHandlers.CashbackList)
	group.PUT("", cashbackHandlers.CashbackUpdate)
	group.DELETE("/:id", cashbackHandlers.CashbackDelete)
	group.PUT("/pass", cashbackHandlers.CashbackPass)
}

func adminInviteGroup(group *gin.RouterGroup) {
	group.GET("", inviteHandlers.InviteRecords)
	group.PUT("", inviteHandlers.InviteUpdate)
	group.PUT("/pass", inviteHandlers.InvitePass)
	group.DELETE("/:id", inviteHandlers.InviteDelete)
}

func adminNotificationGroup(group *gin.RouterGroup) {
	group.GET("", notificationHandlers.List)
	group.POST("", notificationHandlers.Add)
	group.PUT("", notificationHandlers.Update)
	group.DELETE("/:id", notificationHandlers.Delete)
}

func adminOrderGroup(group *gin.RouterGroup) {
	group.GET("", orderHandlers.OrderList)
}

func adminPayGroup(group *gin.RouterGroup) {
	group.GET("", payHandlers.PayConfigList)
	group.POST("", payHandlers.PayConfigAdd)
	group.PUT("", payHandlers.PayConfigUpdate)
	group.DELETE("/:id", payHandlers.PayConfigDelete)
}

func adminConfigGroup(group *gin.RouterGroup) {
	group.GET("", configHandlers.ConfigList)
	group.PUT("", configHandlers.ConfigUpdate)
}

func adminTokenGroup(group *gin.RouterGroup) {
	group.GET("", tokenHandlers.TokenList)
	group.POST("", tokenHandlers.TokenAdd)
	group.PUT("", tokenHandlers.TokenUpdate)
	group.DELETE("/:id", tokenHandlers.TokenDelete)
	// todo
	group.POST("/check", tokenHandlers.TokenCheck)
}

func adminTurnoverGroup(group *gin.RouterGroup) {
	group.GET("", turnoverHandlers.TurnoverList)
	group.DELETE("/:id", turnoverHandlers.TurnoverDelete)
	group.PUT("", turnoverHandlers.TurnoverUpdate)
}

func adminSigninGroup(group *gin.RouterGroup) {
	group.GET("", signinHandlers.SigninList)
}

func adminMessageGroup(group *gin.RouterGroup) {
	group.GET("", messageHandlers.MessageList)
}

func adminProductGroup(group *gin.RouterGroup) {
	group.GET("", productHandlers.ProductList)
	group.POST("", productHandlers.ProductAdd)
	group.PUT("", productHandlers.ProductPut)
	group.DELETE("/:id", productHandlers.ProductDelete)
}

func adminUserGroup(group *gin.RouterGroup) {
	group.GET("", userHandlers.UserList)
	group.DELETE("/:id", userHandlers.UserDelete)
	group.PUT("", userHandlers.UserUpdate)
}

func adminCarmiGroup(group *gin.RouterGroup) {
	group.GET("", carmiHandlers.CarmiList)
	group.POST("", carmiHandlers.CarmiGen)
	group.DELETE("/:id", carmiHandlers.CarmiDel)
	// TODO
	group.GET("/check", carmiHandlers.CarmiCheck)
}
