package gpt

import (
	"chatgpt-web-new-go/common/aiClient"
	"chatgpt-web-new-go/common/config"
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

	// aiClient init
	aiClient.Init()

	code := m.Run()
	os.Exit(code)
}

func TestListModels(t *testing.T) {
	models, err := aiClient.GetGptClient().OpenAIClient.ListModels(context.Background())
	fmt.Println(err)
	fmt.Println(models)

	for _, m := range models.Models {
		fmt.Println(m)
	}
}
