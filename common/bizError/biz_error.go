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

	CarmiStatusError = NewBizError(130000, "卡密状态异常！")
	CarmiUseError    = NewBizError(130001, "卡密使用异常！")
	CarmiDelError    = NewBizError(130003, "卡密删除异常！")

	UserDelError    = NewBizError(140000, "用户删除失败！")
	UserUpdateError = NewBizError(140001, "用户更新失败！")

	ProductUpdateError = NewBizError(150000, "商品更新异常！")
	ProductDeleteError = NewBizError(150001, "商品删除异常！")

	TurnoverDeleteError = NewBizError(160000, "消费记录删除异常！")
	TurnoverUpdateError = NewBizError(160001, "消费记录更新异常！")

	AiKeyTokenDeleteError = NewBizError(170000, "Token删除异常！")
	AiKeyTokenUpdateError = NewBizError(170001, "Token更新异常！")
)
