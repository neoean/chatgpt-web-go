package base

type PageResponse struct {
	Count int64       `json:"count"`
	Rows  interface{} `json:"rows"` // slice
}

type UpdateResponse []interface{}
