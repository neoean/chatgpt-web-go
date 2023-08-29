package user

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(ctx *gin.Context, account string) (*model.User, error) {
	du := dao.Q.User
	userDao := du.WithContext(ctx)

	u, err := userDao.Where(du.Account.Eq(account)).First()
	if err != nil {
		logs.Error("userDao get error: %v", err)
		return nil, err
	}

	return u, err
}
