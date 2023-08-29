package base

type Page struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"size" form:"page_size"`
}
