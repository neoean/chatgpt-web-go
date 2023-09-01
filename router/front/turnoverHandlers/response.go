package turnoverHandlers

import "chatgpt-web-new-go/model"

type TurnoverListData struct {
	Count int64             `json:"count"`
	Rows  []*model.Turnover `json:"rows"`
}
