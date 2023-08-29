package userHandlers

type UserResponse struct {
	Avatar      string `json:"avatar"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Id          int64  `json:"id"`
	Type        int    `json:"type"`
	ReferNumber int64  `json:"refer_number"`
}
