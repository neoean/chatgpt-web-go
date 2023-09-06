package gpt

import (
	"chatgpt-web-new-go/common/config"
	gogpt "github.com/sashabaranov/go-openai"
	"math/rand"
)

func Client() *gogpt.Client {
	keyCount := len(config.Gpt)

	index := rand.Intn(keyCount)

	return config.Gpt[index]
}
