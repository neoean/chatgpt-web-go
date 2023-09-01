package carmi

import (
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

const (
	CarmiStatusCanuse  = 0
	CarmiStatusUsed    = 1
	CarmiStatusExpired = 2
)

func UseCarmi(ctx context.Context, uid int64, carmi, ip string) error {
	dc := dao.Q.Carmi

	c, err := dc.WithContext(ctx).Where(dc.Key.Eq(carmi)).First()
	if err != nil {
		logs.Error("carmiHandlers find carmiHandlers: %v error: %v", carmi, err)
		return err
	}

	if c.Status != CarmiStatusCanuse {
		return bizError.CarmiStatusError
	}

	// update user
	nu := CarmiToUser(ctx, uid, c)
	du := dao.Q.User
	updateUserInfo, err := du.WithContext(ctx).Where(du.ID.Eq(uid)).Updates(nu)
	if err != nil {
		logs.Error("user update error: %v", err)
		return err
	}
	if updateUserInfo.RowsAffected < 1 {
		logs.Error("user update fail RowsAffected < 1: %v", updateUserInfo)
		return err
	}

	// update carmiHandlers
	nc := model.Carmi{
		IP:     ip,
		UserID: uid,
		Status: CarmiStatusUsed,
	}
	updates, err := dc.WithContext(ctx).Where(dc.ID.Eq(c.ID)).Updates(nc)
	if err != nil {
		logs.Error("carmiHandlers use error: %v", err)
		return err
	}
	if updates.RowsAffected < 1 {
		logs.Error("carmiHandlers use fail updates.RowsAffected < 1 : %v", updates)
		return bizError.CarmiUseError
	}

	return nil
}
