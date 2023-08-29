package auth

import "github.com/gin-gonic/gin"

func GetClientIP(c *gin.Context) string {
	clientIP := c.ClientIP()
	if clientIP == "" {
		clientIP = c.RemoteIP()
	}

	return clientIP
}

func GetUA(c *gin.Context) string {
	return c.GetHeader("User-Agent")
}
