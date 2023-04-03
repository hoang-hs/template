package middlewares

import (
	"github.com/gin-gonic/gin"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

// AuthenticateServer authenticate request from server
func (a *AuthService) AuthenticateServer() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		return
	}
}
