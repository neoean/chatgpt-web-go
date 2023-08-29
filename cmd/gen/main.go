package main

import (
	"chatgpt-web-new-go/common/config"
	"chatgpt-web-new-go/common/db"
	"chatgpt-web-new-go/common/logs"
)

func main() {
	// config init
	config.InitConfig()

	// log init
	logs.LogInit()

	// db init
	db.Init()

	// gen
	db.InitGen()
}
