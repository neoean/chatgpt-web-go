package authHandlers

type CodeSendRequest struct {
	Account    string `json:"account"`
	Code       string `json:"code"`
	InviteCode string `json:"invite_code"`
	Password   string `json:"password"`
}
