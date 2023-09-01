package dialog

import (
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func List(ctx context.Context, page, size int) (list []*model.Dialog, count int64, err error) {
	du := dao.Q.Dialog

	list, count, err = du.WithContext(ctx).Where(du.IsDelete.Eq(0)).FindByPage((page-1)*size, size)
	if err != nil {
		logs.Error("Dialog list error: %v", err)
		return nil, 0, err
	}
	return
}

func Delete(ctx context.Context, id int64) error {
	du := dao.Q.Dialog

	resultInfo, err := du.WithContext(ctx).Where(du.ID.Eq(id)).Update(du.IsDelete, 1)
	if err != nil {
		logs.Error("Dialog delete error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("Dialog delete fail: RowsAffected < 1")
		return bizError.CommonDeleteError
	}

	return nil
}

func Add(ctx context.Context, n *model.Dialog) error {
	dn := dao.Q.Dialog

	err := dn.WithContext(ctx).Create(n)
	if err != nil {
		logs.Error("Dialog create error: %v", err)
		return err
	}

	return nil
}
func Update(ctx context.Context, u *model.Dialog) error {
	du := dao.Q.Dialog
	resultInfo, err := du.WithContext(ctx).Where(du.ID.Eq(u.ID)).Updates(u)
	if err != nil {
		logs.Error("Dialog update error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("Dialog update fail: RowsAffected < 1")
		return bizError.CommonUpdateError
	}

	return nil
}
