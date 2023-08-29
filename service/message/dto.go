package message

type Options struct {
	FrequencyPenalty int32  `json:"frequency_penalty"`
	MaxTokens        int32  `json:"max_tokens"`
	Model            string `json:"model"`
	PresencePenalty  int32  `json:"presence_penalty"`
	Temperature      int32  `json:"temperature"`
}

type RequestOptions struct {
	ParentMessageID string   `json:"parentMessageId"`
	Prompt          string   `json:"prompt"`
	Options         *Options `json:"options"`
}

type MessageData struct {
	UserID         int64           `json:"user_id"`
	DateTime       string          `json:"dateTime"`
	Role           string          `json:"role"`
	Status         string          `json:"status"`
	Text           string          `json:"text"`
	PersonaID      int64           `json:"persona_id"`
	PluginID       int64           `json:"plugin_id"`
	PluginInfo     interface{}     `json:"plugin_info"`
	RequestOptions *RequestOptions `json:"requestOptions"`
}

type Message struct {
	Data      []*MessageData `json:"data"`
	ID        string         `json:"id"`
	Path      string         `json:"path"`
	Name      string         `json:"name"`
	PersonaID int64          `json:"persona_id"`
}
