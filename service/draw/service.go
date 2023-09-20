package draw

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

const (
	drawListTypeMe      = "me"
	drawListTypeGallery = "gallery"
)

func DrawList(ctx context.Context, uid int64, _type string, page, size int) (drawRecords []*model.DrawRecord, count int64, err error) {
	do := dao.Q.DrawRecord

	if _type == drawListTypeMe {
		drawRecords, count, err = do.WithContext(ctx).Where(do.IsDelete.Eq(0), do.UserID.Eq(uid)).FindByPage((page-1)*size, size)
	} else {
		drawRecords, count, err = do.WithContext(ctx).Where(do.IsDelete.Eq(0)).FindByPage((page-1)*size, size)
	}
	if err != nil {
		logs.Error("draw list error: %v", err)
		return nil, 0, err
	}
	return
}
