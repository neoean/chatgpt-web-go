package carmi

import (
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/common/random"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func CarmiList(ctx context.Context, page, size int) (carmiList []*model.Carmi, count int64, err error) {
	dc := dao.Q.Carmi

	carmiList, count, err = dc.WithContext(ctx).Where(dc.IsDelete.Eq(0)).FindByPage((page-1)*size, size)
	if err != nil {
		logs.Error("carmi list error: %v", err)
		return nil, 0, err
	}

	return
}

func CarmiDel(ctx context.Context, id int64) error {
	dc := dao.Q.Carmi

	resultInfo, err := dc.WithContext(ctx).Where(dc.ID.Eq(id)).UpdateSimple(dc.IsDelete.Value(1))
	if err != nil {
		logs.Error("carmi update is_delete error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("carmi update is_delete error: RowsAffected < 1")
		return bizError.CarmiDelError
	}

	return nil
}

func CarmiGen(ctx context.Context, r *CarmiGenRequest) (result []*model.Carmi, err error) {
	var carmiList []*model.Carmi
	for i := 0; i < r.Quantity; i++ {
		c := &model.Carmi{
			Key:     random.GenCarmiKey(),
			Value:   int32(r.Reward),
			Type:    r.Type,
			EndTime: r.EndTime,
		}
		carmiList = append(carmiList, c)
	}

	dc := dao.Q.Carmi
	err = dc.WithContext(ctx).CreateInBatches(carmiList, 50)
	if err != nil {
		logs.Error("carmi batch insert error: %v", err)
		return nil, err
	}
	return
}
