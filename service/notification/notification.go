package notification

import (
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func GetList(ctx context.Context, page, size int) (notifications []*model.Notification, count int64, err error) {
	dn := dao.Q.Notification
	notifications, count, err = dn.WithContext(ctx).Where(dn.IsDelete.Eq(0)).FindByPage((page-1)*size, size)
	if err != nil {
		logs.Error("notifications list error: %v", err)
		return nil, 0, err
	}
	return
}

func Add(ctx context.Context, n *model.Notification) error {
	dn := dao.Q.Notification

	err := dn.WithContext(ctx).Create(n)
	if err != nil {
		logs.Error("notifications create error: %v", err)
		return err
	}

	return nil
}

func Update(ctx context.Context, n *model.Notification) error {
	dn := dao.Q.Notification

	resultInfo, err := dn.WithContext(ctx).Where(dn.ID.Eq(n.ID)).Updates(n)
	if err != nil {
		logs.Error("notifications update error: %v", err)
		return err
	}

	if resultInfo.RowsAffected < 1 {
		logs.Error("notification update fail: RowsAffected < 1")
		return bizError.NotificationUpdateError
	}

	return nil
}

func Delete(ctx context.Context, id int64) error {
	dn := dao.Q.Notification

	resultInfo, err := dn.WithContext(ctx).Where(dn.ID.Eq(id)).Update(dn.IsDelete, 1)
	if err != nil {
		logs.Error("notifications delete error: %v", err)
		return err
	}

	if resultInfo.RowsAffected < 1 {
		logs.Error("notification delete fail: RowsAffected < 1")
		return bizError.NotificationDeleteError
	}

	return nil
}
