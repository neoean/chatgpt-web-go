package productHandlers

import "chatgpt-web-new-go/model"

type ProductListData struct {
	Products []*model.Product `json:"products"`
	PayTypes []string         `json:"pay_types"`
}
