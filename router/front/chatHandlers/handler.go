package chatHandlers

import (
	authGlobal "chatgpt-web-new-go/common/auth"
	"chatgpt-web-new-go/common/goUtil"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/common/types"
	"chatgpt-web-new-go/model"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/gpt"
	"chatgpt-web-new-go/service/message"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

func ChatListHandler(c *gin.Context) {

}

func ChatCompletionsHandler(c *gin.Context) {
	var request gpt.Request
	err := c.BindJSON(&request)
	if err != nil {
		base.Fail(c, err.Error())
		return
	}
	logs.Info("chat process request: %v", request)

	if len(request.Prompt) == 0 {
		base.Fail(c, "request messages required")
		return
	}

	uFromCtx, found := c.Get(authGlobal.GinCtxKey)
	if !found {
		logs.Error("cannot get auth user key:%v", authGlobal.GinCtxKey)
		base.Fail(c, "cannot get auth user key:"+authGlobal.GinCtxKey)
		return
	}

	u := uFromCtx.(*model.User)

	stream, err := gpt.Process(c, &request, u.ID)
	if err != nil {
		logs.Error("aiClient Process bizError: %v", err)
		base.Fail(c, "错误:"+err.Error())
		return
	}

	defer stream.Close()

	chanStream := make(chan *message.ChatProcessResponse, 100)
	go func() {
		defer stream.Close()
		defer close(chanStream)
		for i := 0; ; i++ {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				logs.Debug("Stream finished")
				chanStream <- message.StopResponse
				return
			}

			if err != nil {
				logs.Error("Stream error: %v\n", err)
				chanStream <- message.ErrorResponse
				return
			}
			if len(response.Choices) == 0 {
				logs.Debug("Stream finished")
				chanStream <- message.StopResponse
				return
			}

			choice := response.Choices[0]
			data := &message.ChatProcessResponse{
				ID:              response.ID,
				Role:            choice.Delta.Role,
				Segment:         message.SegmentText,
				DateTime:        time.Now().Format(types.TimeFormatDate),
				Content:         choice.Delta.Content,
				ParentMessageID: message.AssistantMessageId,
			}

			if i == 0 {
				data.Segment = message.SegmentStart
			}

			if choice.FinishReason == message.SegmentStop {
				data.Segment = message.SegmentStop
			}

			chanStream <- data

			if choice.FinishReason == message.SegmentStop {
				return
			}
		}
	}()

	var msgList []*message.ChatProcessResponse
	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-chanStream; ok {
			if msg == message.ErrorResponse {
				return false
			}

			msgList = append(msgList, msg)

			marshal, _ := json.Marshal(msg)
			write, err := w.Write(marshal)
			if err != nil {
				logs.Error("w write error: %v, write: %v", err, write)
				return false
			}
			if msg == message.StopResponse {
				return false
			}

			_, _ = w.Write([]byte("\n\n"))
			return true
		}
		return false
	})

	goUtil.New(func() {
		message.MessageAdd(c, u.ID, &request, msgList)
	})
}
