package authHandlers

import (
	"chatgpt-web-new-go/model"
)

const (
	ModelOfficial   = "ChatGPTAPI"
	ModelUnofficial = "ChatGPTUnofficialProxyAPI"
)

type UserInfo struct {
	*model.User
	IsSignIn int `json:"is_signin"`
}

type LoginResponse struct {
	UserInfo UserInfo `json:"user_info"`
	Token    string   `json:"token"`
}

type SessionResponse struct {
	Auth  bool   `json:"auth"`
	Model string `json:"model"`
}
