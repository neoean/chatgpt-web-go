package persona

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func PersonaList(ctx context.Context) (personaList []*model.Persona, err error) {
	dp := dao.Q.Persona

	personaList, err = dp.WithContext(ctx).Find()
	if err != nil {
		logs.Error("persona find error: %v", err)
		return
	}

	return
}
