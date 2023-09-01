package product

import (
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func ProductList(ctx context.Context, page, size int) (productList []*model.Product, count int64, err error) {
	dp := dao.Q.Product

	productList, count, err = dp.WithContext(ctx).Where(dp.IsDelete.Eq(0)).FindByPage((page-1)*size, size)
	if err != nil {
		logs.Error("product find by page error: %v", err)
		return nil, 0, err
	}

	return
}

func ProductAdd(ctx context.Context, p *model.Product) (result *model.Product, err error) {
	dp := dao.Q.Product
	err = dp.WithContext(ctx).Create(p)
	if err != nil {
		logs.Error("product create error: %v", err)
		return nil, err
	}
	return
}

func ProductUpdate(ctx context.Context, p *model.Product) error {
	dp := dao.Q.Product
	resultInfo, err := dp.WithContext(ctx).Where(dp.ID.Eq(p.ID)).Updates(p)
	if err != nil {
		logs.Error("product update error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("product update fail: RowsAffected < 1")
		return bizError.ProductUpdateError
	}

	return nil
}

func ProductDelete(ctx context.Context, id int64) error {
	dp := dao.Q.Product

	resultInfo, err := dp.WithContext(ctx).Where(dp.ID.Eq(id)).Update(dp.ID, 1)
	if err != nil {
		logs.Error("product delete error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("product delete fail: RowsAffected < 1")
		return bizError.ProductDeleteError
	}

	return nil
}
