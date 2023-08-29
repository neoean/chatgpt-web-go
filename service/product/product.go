package product

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func ProductList(ctx context.Context, page, size int) (productList []*model.Product, count int64, err error) {
	dp := dao.Q.Product
	pDao := dp.WithContext(ctx)

	productList, count, err = pDao.FindByPage((page-1)*size, size)
	if err != nil {
		logs.Error("product find by page error: %v", err)
		return nil, 0, err
	}

	return
}
