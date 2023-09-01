package auth

import (
	"chatgpt-web-new-go/common/auth"
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/config"
	"chatgpt-web-new-go/common/constant"
	"chatgpt-web-new-go/common/inviteCodeGen"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/common/redis"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(ctx *gin.Context, account, password, code, inviteCode string) (*model.User, string, error) {
	// 0. valid
	if password == "" && code == "" {
		return nil, "", bizError.LoginPassCodeNoneError
	}

	// 1. login or register
	du := dao.Q.User
	userDao := du.WithContext(ctx)

	u, err := userDao.Where(du.Account.Eq(account)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logs.Error("user FirstOrCreate by account:%v bizError:%v", account, err)
		return nil, "", err
	}

	// register or login
	if u == nil {
		return doRegister(ctx, account, password, code, inviteCode)
	} else {
		return doLogin(u, account, password, code)
	}
}

func doLogin(u *model.User, account, password, code string) (*model.User, string, error) {
	// password valid
	if password != "" && !u.ComparePassword(password) {
		return nil, "", bizError.LoginPasswordError
	}
	if code != "" && !loginCodeValid(account, code) {
		return nil, "", bizError.LoginCodeErrorError
	}

	// session key token
	token, err := auth.Encode(u)
	if err != nil {
		logs.Error("user auth encode phone:%v bizError:%v", account, err)
		return nil, "", err
	}
	return u, token, nil
}

func doRegister(ctx *gin.Context, account, password, code, inviteCode string) (*model.User, string, error) {
	// 0. valid
	if password == "" {
		return nil, "", bizError.LoginPassCodeNoneError
	}

	pass := loginCodeValid(account, code)
	if !pass {
		return nil, "", bizError.LoginCodeErrorError
	}

	// 1. user model
	u := &model.User{
		Account:   account,
		Password:  password,
		IP:        auth.GetClientIP(ctx),
		UserAgent: auth.GetUA(ctx),
	}

	// 2. get super id
	du := dao.Q.User
	su, err := du.WithContext(ctx).Where(du.InviteCode.Eq(inviteCode)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logs.Error("user get by inviteCode error: %v", err)
		return nil, "", err
	}
	if su != nil {
		u.SuperiorID = su.ID
	}

	err = du.WithContext(ctx).Create(u)
	if err != nil {
		logs.Error("user Create by account:%v bizError:%v", account, err)
		return nil, "", err
	}

	// update invite code
	ic := inviteCodeGen.InviteCodeGen.IdToCode(u.ID)
	updateColumn, err := du.WithContext(ctx).Where(du.ID.Eq(u.ID)).UpdateColumn(du.InviteCode, ic)
	if err != nil {
		logs.Error("user update by account:%v bizError:%v", account, err)
		return nil, "", err
	}
	logs.Info("user invite code update result: %v", updateColumn)

	// 3. session key token
	token, err := auth.Encode(u)
	if err != nil {
		logs.Error("user auth encode phone:%v bizError:%v", account, err)
		return nil, "", err
	}

	return u, token, nil
}

// only one time
func loginCodeValid(phone string, code string) bool {
	if _, found := constant.WhiteListPhone[phone]; found {
		return true
	}

	key := fmt.Sprintf(redis.KeySnsCode, phone)
	codeFromRedis, err := config.Redis.Get(key).Result()
	if err != nil {
		logs.Warn("redis get: %v bizError: %v", key, err)
		return false
	}

	if code != codeFromRedis {
		logs.Info("code:%v != codeFromRedis:%v", code, codeFromRedis)
		return false
	}

	// expire code
	_ = config.Redis.Del(key)

	return true
}
