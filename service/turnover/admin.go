package turnover

import (
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func AdminTurnoverList(ctx context.Context, page, size int) (turnoverList []*model.Turnover, count int64, err error) {
	dt := dao.Q.Turnover
	turnoverList, count, err = dt.WithContext(ctx).Where(dt.IsDelete.Eq(0)).FindByPage((page-1)*size, size)
	if err != nil {
		logs.Error("turnover list error: %v", err)
		return nil, 0, err
	}

	return
}

func AdminTurnoverDelete(ctx context.Context, id int64) error {
	dt := dao.Q.Turnover
	resultInfo, err := dt.WithContext(ctx).Where(dt.ID.Eq(id)).Update(dt.IsDelete, 1)
	if err != nil {
		logs.Error("turnover delete error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("turnover delete fail: RowsAffected < 1")
		return bizError.TurnoverDeleteError
	}

	return nil
}

func AdminTurnoverUpdate(ctx context.Context, t *model.Turnover) (result *model.Turnover, err error) {
	dt := dao.Q.Turnover
	resultInfo, err := dt.WithContext(ctx).Where(dt.ID.Eq(t.ID)).Updates(t)
	if err != nil {
		logs.Error("turnover update error: %v", err)
		return nil, err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("turnover update fail: RowsAffected < 1 : %v", resultInfo)
		return nil, bizError.TurnoverUpdateError
	}

	return
}
