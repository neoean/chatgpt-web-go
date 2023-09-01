package withdrawal

import (
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func List(ctx context.Context, page, size int) (list []*model.WithdrawalRecord, count int64, err error) {
	du := dao.Q.WithdrawalRecord

	list, count, err = du.WithContext(ctx).Where(du.IsDelete.Eq(0)).FindByPage((page-1)*size, size)
	if err != nil {
		logs.Error("WithdrawalRecord list error: %v", err)
		return nil, 0, err
	}
	return
}

func Delete(ctx context.Context, id int64) error {
	du := dao.Q.WithdrawalRecord

	resultInfo, err := du.WithContext(ctx).Where(du.ID.Eq(id)).Update(du.IsDelete, 1)
	if err != nil {
		logs.Error("WithdrawalRecord delete error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("WithdrawalRecord delete fail: RowsAffected < 1")
		return bizError.CommonDeleteError
	}

	return nil
}

func Update(ctx context.Context, u *model.WithdrawalRecord) error {
	du := dao.Q.WithdrawalRecord
	resultInfo, err := du.WithContext(ctx).Where(du.ID.Eq(u.ID)).Updates(u)
	if err != nil {
		logs.Error("WithdrawalRecord update error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("WithdrawalRecord update fail: RowsAffected < 1")
		return bizError.CommonUpdateError
	}

	return nil
}
