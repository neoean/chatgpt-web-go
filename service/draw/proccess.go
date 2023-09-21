package draw

import (
	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-openai"
)

func Process(ctx *gin.Context, r *Request, uid int64) (stream *gogpt.ChatCompletionStream, err error) {

	return
}
