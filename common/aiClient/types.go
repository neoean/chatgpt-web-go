package aiClient

import (
	"chatgpt-web-new-go/model"
	sd "github.com/SpenserCai/sd-webui-go"
	"github.com/sashabaranov/go-openai"
)

var (
	GptClients    []*GptClient
	DallE2Clients []*DallE2Client
	SDClients     []*SDClient
)

type GptClient struct {
	OpenAIClient *openai.Client
	Model        *model.Aikey
}

type DallE2Client struct {
	OpenAIClient *openai.Client
	Model        *model.Aikey
}

type SDClient struct {
	SdClient sd.StableDiffInterface
	Model    *model.Aikey
}
