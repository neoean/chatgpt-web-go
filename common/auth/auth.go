package auth

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/model"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

var (
	GinCtxKey = "authUser"
	key       = []byte("jansdfjizxuqhawiwehjioas7812738_asdf+787")
)

type CustomClaims struct {
	User *model.User
	jwt.StandardClaims
}

// Decode a token string into a token object
func Decode(tokenString string) (*CustomClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		logs.Error("Decode jwt.ParseWithClaims errorL %v", err)
		return nil, err
	}

	// Validate the token and return the custom claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		logs.Error("Decode token.Claims errorL %v", err)
		return nil, err
	}
}

// Encode a claim into a JWT
func Encode(user *model.User) (string, error) {
	expireToken := time.Now().Add(time.Hour * 7200).Unix()

	// Create the Claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "chatgpt-web-new-go",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)
}

// EncodeByCtx 从ctx中的token获取登录用户信息
func EncodeByCtx(c *gin.Context) (*CustomClaims, error) {
	//1.获取token
	token := c.GetHeader("Token")
	if token != "" {
		tokenS := strings.Split(token, " ")
		token = tokenS[0]
	} else {
		token = c.Request.FormValue("token")
	}
	if token == "" || token == "undefined" {
		return nil, errors.New("not found token")
	}

	return Decode(token)
}
