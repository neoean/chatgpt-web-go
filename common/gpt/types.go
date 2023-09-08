package gpt

import (
	"chatgpt-web-new-go/model"
	gogpt "github.com/sashabaranov/go-openai"
)

var Clients []*GptClient

type GptClient struct {
	GoGptClient *gogpt.Client
	Model       *model.Aikey
}
