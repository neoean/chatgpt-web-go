package turnover

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func TurnoverList(ctx context.Context, uid int64, page, size int) (turnovers []*model.Turnover, count int64, err error) {
	dt := dao.Q.Turnover
	tDao := dt.WithContext(ctx)

	offset := (page - 1) * size
	turnovers, count, err = tDao.Where(dt.UserID.Eq(uid)).FindByPage(offset, size)
	if err != nil {
		logs.Error("turnover find by page error: %v, uid: %v", err, uid)
		return nil, 0, err
	}

	return
}
