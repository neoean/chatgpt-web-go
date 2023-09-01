package carmi

import (
	"chatgpt-web-new-go/model"
	"time"
)

const (
	CarmiTypeIntegral = "integral"
	CarmiTypeVipDays  = "day"

	CarmiLevelVip  = 1
	CarmiLevelSVip = 2
)

func CarmiToUser(uid int64, carmi *model.Carmi) (u *model.User) {
	u = &model.User{
		ID: uid,
	}

	if carmi.Type == CarmiTypeIntegral {
		u.Integral = carmi.Value
	}

	if carmi.Type == CarmiTypeVipDays {
		et := time.Now().AddDate(0, 0, int(carmi.Value))
		if carmi.Level == CarmiLevelVip {
			u.VipExpireTime = et
		}
		if carmi.Level == CarmiLevelSVip {
			u.SvipExpireTime = et
		}
	}

	return u
}
