package invite

import (
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

const (
	InviteStatusFail    = 0
	InviteStatusPass    = 1
	InviteStatusPending = 3
)

func InviteRecords(ctx context.Context, page, size int) (inviteRecords []*model.InviteRecord, count int64, err error) {
	di := dao.Q.InviteRecord

	inviteRecords, count, err = di.WithContext(ctx).Where(di.IsDelete.Eq(0)).FindByPage((page-1)*size, size)
	if err != nil {
		logs.Error("invite records error: %v", err)
		return nil, 0, err
	}
	return
}

func InviteUpdate(ctx context.Context, invite *model.InviteRecord) error {
	di := dao.Q.InviteRecord
	resultInfo, err := di.WithContext(ctx).Updates(invite)
	if err != nil {
		logs.Error("invite update error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("invite update fail: RowsAffected < 1")
		return bizError.CommonUpdateError
	}

	return nil
}

func InviteDelete(ctx context.Context, id int64) error {
	di := dao.Q.InviteRecord
	resultInfo, err := di.WithContext(ctx).Where(di.ID.Eq(id)).Update(di.IsDelete, 1)
	if err != nil {
		logs.Error("invite delete error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("invite delete fail: RowsAffected < 1")
		return bizError.CommonUpdateError
	}

	return nil
}

func InvitePass(ctx context.Context) error {
	di := dao.Q.InviteRecord

	_, err := di.WithContext(ctx).Where(di.Status.Eq(InviteStatusPending)).Update(di.Status, InviteStatusPass)
	if err != nil {
		logs.Error("invite pass error: %v", err)
		return err
	}
	return nil
}
