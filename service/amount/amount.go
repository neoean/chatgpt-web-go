package amount

import (
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func List(ctx context.Context, page, size int) (list []*model.AmountDetail, count int64, err error) {
	du := dao.Q.AmountDetail

	list, count, err = du.WithContext(ctx).Where(du.IsDelete.Eq(0)).FindByPage((page-1)*size, size)
	if err != nil {
		logs.Error("AmountDetail list error: %v", err)
		return nil, 0, err
	}
	return
}

func Delete(ctx context.Context, id int64) error {
	du := dao.Q.AmountDetail

	resultInfo, err := du.WithContext(ctx).Where(du.ID.Eq(id)).Update(du.IsDelete, 1)
	if err != nil {
		logs.Error("AmountDetail delete error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("AmountDetail delete fail: RowsAffected < 1")
		return bizError.CommonDeleteError
	}

	return nil
}

func Update(ctx context.Context, u *model.AmountDetail) error {
	du := dao.Q.AmountDetail
	resultInfo, err := du.WithContext(ctx).Where(du.ID.Eq(u.ID)).Updates(u)
	if err != nil {
		logs.Error("AmountDetail update error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("AmountDetail update fail: RowsAffected < 1")
		return bizError.CommonUpdateError
	}

	return nil
}
