package plugin

import (
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func List(ctx context.Context, page, size int) (list []*model.Plugin, count int64, err error) {
	du := dao.Q.Plugin

	list, count, err = du.WithContext(ctx).Where(du.IsDelete.Eq(0)).FindByPage((page-1)*size, size)
	if err != nil {
		logs.Error("Plugin list error: %v", err)
		return nil, 0, err
	}
	return
}

func Delete(ctx context.Context, id int64) error {
	du := dao.Q.Plugin

	resultInfo, err := du.WithContext(ctx).Where(du.ID.Eq(id)).Update(du.IsDelete, 1)
	if err != nil {
		logs.Error("Plugin delete error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("Plugin delete fail: RowsAffected < 1")
		return bizError.CommonDeleteError
	}

	return nil
}

func Add(ctx context.Context, n *model.Plugin) error {
	dn := dao.Q.Plugin

	err := dn.WithContext(ctx).Create(n)
	if err != nil {
		logs.Error("Plugin create error: %v", err)
		return err
	}

	return nil
}

func Update(ctx context.Context, u *model.Plugin) error {
	du := dao.Q.Plugin
	resultInfo, err := du.WithContext(ctx).Where(du.ID.Eq(u.ID)).Updates(u)
	if err != nil {
		logs.Error("Plugin update error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("Plugin update fail: RowsAffected < 1")
		return bizError.CommonUpdateError
	}

	return nil
}
