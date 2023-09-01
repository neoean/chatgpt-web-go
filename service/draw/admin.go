package draw

import (
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func List(ctx context.Context, page, size int) (list []*model.DrawRecord, count int64, err error) {
	du := dao.Q.DrawRecord

	list, count, err = du.WithContext(ctx).Where(du.IsDelete.Eq(0)).FindByPage((page-1)*size, size)
	if err != nil {
		logs.Error("DrawRecord list error: %v", err)
		return nil, 0, err
	}
	return
}

func Delete(ctx context.Context, id int64) error {
	du := dao.Q.DrawRecord

	resultInfo, err := du.WithContext(ctx).Where(du.ID.Eq(id)).Update(du.IsDelete, 1)
	if err != nil {
		logs.Error("DrawRecord delete error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("DrawRecord delete fail: RowsAffected < 1")
		return bizError.CommonDeleteError
	}

	return nil
}

func Add(ctx context.Context, n *model.DrawRecord) error {
	dn := dao.Q.DrawRecord

	err := dn.WithContext(ctx).Create(n)
	if err != nil {
		logs.Error("DrawRecord create error: %v", err)
		return err
	}

	return nil
}

func Update(ctx context.Context, u *model.DrawRecord) error {
	du := dao.Q.DrawRecord
	resultInfo, err := du.WithContext(ctx).Where(du.ID.Eq(u.ID)).Updates(u)
	if err != nil {
		logs.Error("DrawRecord update error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("DrawRecord update fail: RowsAffected < 1")
		return bizError.CommonUpdateError
	}

	return nil
}
