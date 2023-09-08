package gpt

import (
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/goUtil"
	"chatgpt-web-new-go/common/gpt"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/common/types"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"fmt"
	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-openai"
)

const (
	coinPerMessage = -1
)

func Process(ctx *gin.Context, r *Request, uid int64) (stream *gogpt.ChatCompletionStream, err error) {
	// Deduction of points
	du := dao.Q.User
	updateInfo, err := du.WithContext(ctx).Where(du.ID.Eq(uid)).UpdateSimple(du.Integral.Add(coinPerMessage))
	if err != nil {
		logs.Error("user integral update error: %v", err)
		return
	}
	if updateInfo.RowsAffected < 1 {
		logs.Error("user integral update error: updateInfo.RowsAffected <= 0")
		err = bizError.IntegralNoneError
		return
	}

	request := gogpt.ChatCompletionRequest{
		Model:            r.Options.Model,
		MaxTokens:        r.Options.MaxTokens,
		Temperature:      r.Options.Temperature,
		PresencePenalty:  r.Options.PresencePenalty,
		FrequencyPenalty: r.Options.FrequencyPenalty,
		Messages: []gogpt.ChatCompletionMessage{
			{
				Role:    gogpt.ChatMessageRoleAssistant,
				Content: r.Prompt,
			},
		},
		Stream: true,
	}

	// cnf.Model 是否在 chatModels 中
	gptClient := gpt.GetClient()
	if gptClient == nil {
		logs.Error("gptClient is nil!")
		return nil, bizError.AiKeyNoneUsefullError
	}

	stream, err = gptClient.GoGptClient.CreateChatCompletionStream(ctx, request)
	if err != nil {
		logs.Error("gpt client.CreateChatCompletion bizError:%v", err)
		goUtil.New(func() {
			refreshKey(gptClient)
		})
		return nil, err
	}

	// insert message
	goUtil.New(func() {
		addMessage(ctx, r, uid)
	})

	// insert record
	goUtil.New(func() {
		insertTurnOverRecord(ctx, r, uid)
	})

	return
}

func refreshKey(client *gpt.GptClient) {
	aiKey := client.Model

	aiKey.Status = 0

	da := dao.Q.Aikey
	resultInfo, err := da.Where(da.ID.Eq(aiKey.ID)).Update(da.Status, 0)
	if err != nil {
		logs.Error("aiKey update error: %v", err)
		return
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("aiKey update fail: RowsAffected < 1")
		return
	}

	// refresh
	gpt.DoInitClient()
}

func insertTurnOverRecord(ctx *gin.Context, r *Request, uid int64) {
	t := &model.Turnover{
		UserID:   uid,
		Describe: "对话(" + r.Options.Model + ")",
		Value:    fmt.Sprintf("%v积分", coinPerMessage),
	}

	dr := dao.Q.Turnover
	err := dr.WithContext(ctx).Create(t)
	if err != nil {
		logs.Error("turnover create error: %v", err)
	}
}

func addMessage(ctx *gin.Context, r *Request, uid int64) {
	msg := &model.Message{
		UserID:           uid,
		Content:          r.Prompt,
		PersonaID:        types.InterfaceToInt64(r.PersonaId),
		Role:             gogpt.ChatMessageRoleUser,
		FrequencyPenalty: int32(r.Options.FrequencyPenalty),
		MaxTokens:        int32(r.Options.MaxTokens),
		Model:            r.Options.Model,
		PresencePenalty:  int32(r.Options.PresencePenalty),
		Temperature:      int32(r.Options.Temperature),
		ParentMessageID:  r.ParentMessageId,
	}

	dm := dao.Q.Message
	err := dm.WithContext(ctx).Create(msg)
	if err != nil {
		logs.Error("message create error: %v", err)
	}
	return
}
