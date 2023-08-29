package password

import (
	"chatgpt-web-new-go/common/logs"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// Hash 进行加密
func Hash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		logs.Error("hash password bizError: %v", err)
	}

	return string(bytes)
}

//CheckHash 检查密码和hash是否匹配
func CheckHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(err)
	return err == nil
}

// IsHashed 检查密码和hash是否已经加密
func IsHashed(str string) bool {
	return len(str) == 60
}
