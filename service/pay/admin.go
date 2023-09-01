package pay

import (
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func PayConfigList(ctx context.Context) (payments []*model.Payment, err error) {
	dp := dao.Q.Payment
	payments, err = dp.WithContext(ctx).Find()
	if err != nil {
		logs.Error("pay config list error: %v", err)
		return nil, err
	}
	return
}

func PayConfigAdd(ctx context.Context, pay *model.Payment) error {
	dp := dao.Q.Payment

	err := dp.WithContext(ctx).Create(pay)
	if err != nil {
		logs.Error("payment create error: %v", err)
		return err
	}

	return nil
}

func PayConfigUpdate(ctx context.Context, pay *model.Payment) error {
	dp := dao.Q.Payment
	resultInfo, err := dp.WithContext(ctx).Where(dp.ID.Eq(pay.ID)).Updates(pay)
	if err != nil {
		logs.Error("payment update error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("payment update fail: RowsAffected < 1")
		return bizError.PaymentUpdateError
	}
	return nil
}

func PayConfigDelete(ctx context.Context, id int64) error {
	dp := dao.Q.Payment

	resultInfo, err := dp.WithContext(ctx).Where(dp.ID.Eq(id)).Update(dp.IsDelete, 1)
	if err != nil {
		logs.Error("payment delete error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("payment delete fail: RowsAffected < 1")
		return bizError.PaymentDeleteError
	}
	return nil
}
