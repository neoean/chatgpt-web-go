package carmi

import (
	"chatgpt-web-new-go/model"
	"context"
	"time"
)

const (
	CarmiTypeIntegral = "integral"
	CarmiTypeVipDays  = "vip_days"
	CarmiTypeSVipDays = "svip_days"
)

var (
	typeToHandler = map[string]func(ctx context.Context, uid int64, carmi *model.Carmi) (u *model.User){
		CarmiTypeIntegral: integralToUser,
		CarmiTypeVipDays:  vipToUser,
		CarmiTypeSVipDays: sVipToUser,
	}
)

func CarmiToUser(ctx context.Context, uid int64, carmi *model.Carmi) (u *model.User) {
	f, _ := typeToHandler[carmi.Type]
	return f(ctx, uid, carmi)
}

func integralToUser(ctx context.Context, uid int64, carmi *model.Carmi) (u *model.User) {
	u = &model.User{
		ID:       uid,
		Integral: carmi.Value,
	}
	return
}

func vipToUser(ctx context.Context, uid int64, carmi *model.Carmi) (u *model.User) {
	et := time.Now().AddDate(0, 0, int(carmi.Value))
	u = &model.User{
		ID:            uid,
		VipExpireTime: et,
	}
	return
}

func sVipToUser(ctx context.Context, uid int64, carmi *model.Carmi) (u *model.User) {
	et := time.Now().AddDate(0, 0, int(carmi.Value))
	u = &model.User{
		ID:             uid,
		SvipExpireTime: et,
	}
	return
}
