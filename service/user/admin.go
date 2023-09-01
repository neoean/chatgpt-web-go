package user

import (
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func UserList(ctx context.Context, page, size int) (userList []*model.User, count int64, err error) {
	du := dao.Q.User

	userList, count, err = du.WithContext(ctx).Where(du.IsDelete.Eq(0)).FindByPage((page-1)*size, size)
	if err != nil {
		logs.Error("user list error: %v", err)
		return nil, 0, err
	}
	return
}

func UserDelete(ctx context.Context, uid int64) error {
	du := dao.Q.User

	resultInfo, err := du.WithContext(ctx).Where(du.ID.Eq(uid)).Update(du.IsDelete, 1)
	if err != nil {
		logs.Error("user delete error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("user delete fail: RowsAffected < 1")
		return bizError.UserDelError
	}

	return nil
}

func UserUpdate(ctx context.Context, u *model.User) error {
	du := dao.Q.User
	resultInfo, err := du.WithContext(ctx).Where(du.ID.Eq(u.ID)).Updates(u)
	if err != nil {
		logs.Error("user update error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("user update fail: RowsAffected < 1")
		return bizError.UserUpdateError
	}

	return nil
}
