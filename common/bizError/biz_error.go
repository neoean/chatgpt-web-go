package bizError

import "fmt"

type BizError struct {
	Code    int    `json:"code"`
	Message string `json:"messageHandlers"`
}

func (e *BizError) Error() string {
	return fmt.Sprintf("[%v]%v", e.Code, e.Message)
}

func NewBizError(code int, message string) *BizError {
	return &BizError{
		Code:    code,
		Message: message,
	}
}

var (
	UnknowError            = NewBizError(-1, "未知错误！")
	LoginCodeNoneError     = NewBizError(100000, "验证码不能为空！")
	LoginPassCodeNoneError = NewBizError(100001, "密码或验证码不能为空！")
	LoginCodeErrorError    = NewBizError(100002, "验证码错误！")
	LoginPasswordError     = NewBizError(100003, "密码错误！")

	SigninedAlreadyError = NewBizError(110000, "已经签到过啦！")

	IntegralNoneError = NewBizError(120000, "积分用光啦！")
)
