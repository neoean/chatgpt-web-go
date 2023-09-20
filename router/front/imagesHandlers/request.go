package imagesHandlers

type ImagesRequest struct {
	IsAll    bool   `json:"isAll"`
	Loading  bool   `json:"loading"`
	Type     string `json:"type"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}
