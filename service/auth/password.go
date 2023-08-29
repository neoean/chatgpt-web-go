package auth

import (
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"context"
)

func UserPasswordReset(ctx context.Context, account, code, password string) error {
	// 0. valid
	if password == "" && code == "" {
		return bizError.LoginPassCodeNoneError
	}

	// 1. valid code
	pass := loginCodeValid(account, code)
	if !pass {
		return bizError.LoginCodeErrorError
	}

	// 2. update
	du := dao.Q.User
	_, err := du.WithContext(ctx).Where(du.Account.Eq(account)).Update(du.Password, pass)
	if err != nil {
		logs.Error("user password update error: %v", err)
		return err
	}

	return nil
}
