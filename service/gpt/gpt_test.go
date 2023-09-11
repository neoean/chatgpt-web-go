package gpt

import (
	"chatgpt-web-new-go/common/config"
	"chatgpt-web-new-go/common/gpt"
	"chatgpt-web-new-go/common/logs"
	"context"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// config init
	config.InitConfig()

	// log init
	logs.LogInit()

	// gpt init
	gpt.Init()

	code := m.Run()
	os.Exit(code)
}

func TestListModels(t *testing.T) {
	models, err := gpt.GetClient().GoGptClient.ListModels(context.Background())
	fmt.Println(err)
	fmt.Println(models)

	for _, m := range models.Models {
		fmt.Println(m)
	}
}
