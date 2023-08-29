package message

import gogpt "github.com/sashabaranov/go-openai"

const (
	SegmentText  = "text"
	SegmentStart = "start"
	SegmentStop  = "stop"
	SegmentError = "error"

	AssistantMessageId = "assistantMessageId"
)

var (
	StartResponse = &ChatProcessResponse{
		Role:            gogpt.ChatMessageRoleAssistant,
		Segment:         SegmentStart,
		ParentMessageID: AssistantMessageId,
	}

	StopResponse = &ChatProcessResponse{
		Role:            gogpt.ChatMessageRoleAssistant,
		Segment:         SegmentStop,
		ParentMessageID: AssistantMessageId,
	}

	ErrorResponse = &ChatProcessResponse{
		Role:            gogpt.ChatMessageRoleAssistant,
		Segment:         SegmentError,
		ParentMessageID: AssistantMessageId,
	}
)

type ChatProcessResponse struct {
	ID              string `json:"id"`
	Role            string `json:"role"`
	Segment         string `json:"segment"`
	DateTime        string `json:"dateTime"`
	Content         string `json:"content"`
	ParentMessageID string `json:"parentMessageId"`
}
