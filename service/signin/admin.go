package signin

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func AdminSigninList(ctx context.Context, page, size int) (signinList []*model.Signin, count int64, err error) {
	ds := dao.Q.Signin
	signinList, count, err = ds.WithContext(ctx).Where(ds.IsDelete.Eq(0)).FindByPage((page-1)*size, size)
	if err != nil {
		logs.Error("signin list error: %v", err)
		return nil, 0, err
	}

	return
}
