package gpt

type Request struct {
	ParentMessageId string      `json:"parentMessageId"`
	PersonaId       interface{} `json:"persona_id"`
	Prompt          string      `json:"prompt"`
	Options         Options     `json:"options"`
}

type Options struct {
	FrequencyPenalty float32 `json:"frequencyPenalty"`
	MaxTokens        int     `json:"max_tokens"`
	Model            string  `json:"model"`
	PresencePenalty  float32 `json:"presence_penalty"`
	Temperature      float32 `json:"temperature"`
}
