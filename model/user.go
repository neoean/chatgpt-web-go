package model

import (
	"chatgpt-web-new-go/common/auth/password"
	"gorm.io/gorm"
)

// ComparePassword 检查密码是否匹配
func (user *User) ComparePassword(_password string) bool {
	return password.CheckHash(_password, user.Password)
}

// BeforeSave 保存前
func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if !password.IsHashed(user.Password) {
		user.Password = password.Hash(user.Password)
	}
	return
}
