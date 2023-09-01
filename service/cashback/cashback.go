package cashback

import (
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

const (
	CashbackStatusFail    = 0
	CashbackStatusPass    = 1
	CashbackStatusPending = 3
)

func CashbackList(ctx context.Context, page, size int) (cashbackList []*model.Cashback, count int64, err error) {
	dc := dao.Q.Cashback

	cashbackList, count, err = dc.WithContext(ctx).Where(dc.IsDelete.Eq(0)).FindByPage((page-1)*size, size)
	if err != nil {
		logs.Error("cashback list error: %v", err)
		return nil, 0, err
	}
	return
}

func CashbackUpdate(ctx context.Context, c *model.Cashback) error {
	dc := dao.Q.Cashback

	resultInfo, err := dc.WithContext(ctx).Updates(c)
	if err != nil {
		logs.Error("cashback update error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("cashback update fail: RowsAffected < 1")
		return bizError.CommonUpdateError
	}
	return nil
}

func CashbackDelete(ctx context.Context, id int64) error {
	dc := dao.Q.Cashback

	resultInfo, err := dc.WithContext(ctx).Where(dc.ID.Eq(id)).Update(dc.IsDelete, 1)
	if err != nil {
		logs.Error("cashback delete error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("cashback delete fail: RowsAffected < 1")
		return bizError.CommonDeleteError
	}
	return nil
}

func CashbackPass(ctx context.Context, id int64) error {
	dc := dao.Q.Cashback

	resultInfo, err := dc.WithContext(ctx).Where(dc.ID.Eq(id)).Update(dc.Status, CashbackStatusPass)
	if err != nil {
		logs.Error("cashback pass error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("cashback pass fail: RowsAffected < 1")
		return bizError.CommonDeleteError
	}
	return nil
}
