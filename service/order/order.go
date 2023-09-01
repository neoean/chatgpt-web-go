package order

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func OrderList(ctx context.Context, page, size int) (orderList []*model.Order, count int64, err error) {
	do := dao.Q.Order
	orderList, count, err = do.WithContext(ctx).Where(do.IsDelete.Eq(0)).FindByPage((page-1)*size, size)
	if err != nil {
		logs.Error("order list error: %v", err)
		return nil, 0, err
	}
	return
}
