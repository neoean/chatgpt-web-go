package token

import (
	"chatgpt-web-new-go/common/aiClient"
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/goUtil"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func TokenList(ctx context.Context, page, size int) (tokenList []*model.Aikey, count int64, err error) {
	dt := dao.Q.Aikey

	tokenList, count, err = dt.WithContext(ctx).Where(dt.IsDelete.Eq(0)).FindByPage((page-1)*size, size)
	if err != nil {
		logs.Error("token list error: %v", err)
		return
	}
	return
}

func TokenAdd(ctx context.Context, token *model.Aikey) (result *model.Aikey, err error) {
	dt := dao.Q.Aikey
	err = dt.WithContext(ctx).Create(token)
	if err != nil {
		logs.Error("token create error: %v", err)
		return nil, err
	}

	goUtil.New(func() {
		aiClient.DoInitClient()
	})

	return
}

func TokenUpdate(ctx context.Context, token *model.Aikey) (result *model.Aikey, err error) {
	dt := dao.Q.Aikey

	resultInfo, err := dt.WithContext(ctx).Where(dt.ID.Eq(token.ID)).Updates(token)
	if err != nil {
		logs.Error("token update error: %v", err)
		return nil, err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("token update fail: RowsAffected < 1")
		return nil, bizError.AiKeyTokenUpdateError
	}

	goUtil.New(func() {
		aiClient.DoInitClient()
	})

	return
}

func TokenDelete(ctx context.Context, id int64) error {
	dt := dao.Q.Aikey

	resultInfo, err := dt.WithContext(ctx).Where(dt.ID.Eq(id)).Update(dt.IsDelete, 1)
	if err != nil {
		logs.Error("token delete error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("token delete fail: RowsAffected < 1")
		return bizError.AiKeyTokenDeleteError
	}

	goUtil.New(func() {
		aiClient.DoInitClient()
	})

	return nil
}
