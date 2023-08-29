package pay

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"context"
)

func PayTypeList(ctx context.Context) (payTypeList []string, err error) {
	dp := dao.Q.Payment
	payDao := dp.WithContext(ctx)

	paymentList, err := payDao.Where(dp.Status.Eq(1)).Find()
	if err != nil {
		logs.Error("pay dao find error: %v", err)
		return nil, err
	}

	for _, p := range paymentList {
		payTypeList = append(payTypeList, p.Types)
	}

	return
}
