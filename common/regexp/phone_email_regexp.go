package regexp

import "regexp"

const (
	// 正则表达式匹配中国合法手机号
	phoneRegex = "^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(17[6|7|8])|(18[0-9])|166|198|199)\\d{8}$"

	// 正则表达式匹配邮箱
	emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

func IsValidPhone(phone string) bool {
	isPhone, _ := regexp.MatchString(phoneRegex, phone)
	return isPhone
}

func isValidEmail(email string) bool {
	isEmail, _ := regexp.MatchString(emailRegex, email)
	return isEmail
}

func IsValidPhoneOrEmail(input string) bool {
	return IsValidPhone(input) || isValidEmail(input)
}
