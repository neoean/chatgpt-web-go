package main

import (
	"chatgpt-web-new-go/common/config"
	"chatgpt-web-new-go/common/db"
	"chatgpt-web-new-go/common/email"
	"chatgpt-web-new-go/common/gpt"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/common/redis"
	"chatgpt-web-new-go/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// config init
	config.InitConfig()

	// log init
	logs.LogInit()

	// db init
	db.Init()

	// redis init
	redis.Init()

	// email service
	email.InitEmailDialer()

	// gpt init
	gpt.Init()

	// gin init
	engine := gin.Default()

	// route init
	router.Init(engine)

	// listen init
	port := fmt.Sprintf("%v", config.Config.Port)
	err := engine.Run("127.0.0.1:" + port)
	if err != nil {
		logs.Error("run webserver error %v", err)
		panic(err)
	}
}
