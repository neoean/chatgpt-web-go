package signin

import (
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/common/types"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
	"fmt"
)

const (
	SigninCoin = 10
)

func SignIn(ctx context.Context, uid int64, ip string) error {
	if IsSignInToday(uid) {
		return bizError.SigninedAlreadyError
	}

	sign := &model.Signin{
		UserID: uid,
		IP:     ip,
	}

	err := dao.Q.Signin.WithContext(ctx).Create(sign)
	if err != nil {
		logs.Error("signin create error: %v", err)
		return err
	}

	// add integral
	du := dao.Q.User
	result, err := du.WithContext(ctx).Where(du.ID.Eq(uid)).UpdateSimple(du.Integral.Add(SigninCoin))
	if err != nil {
		logs.Error("user integral update error: %v", err)
		return err
	}
	if result.RowsAffected < 1 {
		err := fmt.Errorf("user integral update RowsAffected < 1, uid: %v", uid)
		logs.Error("user integral update error: %v", err)
		return err
	}

	return nil
}

func IsSignInToday(uid int64) bool {
	startTime, endTime := types.GetDayStartEn()

	sign := dao.Q.Signin
	signInToday, err := sign.
		Where(sign.UserID.Eq(uid),
			sign.CreateTime.Between(startTime, endTime)).
		Find()
	if err != nil {
		logs.Error("signIn today get error: %v", err)
		return false
	}
	return len(signInToday) > 0
}

func GetSignListMonth(uid int64) (result []*model.Signin, err error) {
	start, end := types.GetMonthStartEnd()

	sign := dao.Q.Signin
	result, err = sign.Where(sign.UserID.Eq(uid), sign.CreateTime.Between(start, end)).Find()

	return
}
