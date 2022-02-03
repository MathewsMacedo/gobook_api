package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var token = "6d6e32f37bd110a2821e65fcb2cb6c63fa1f7e4198cbf488b81506f28ff92350ddede4638e46a8dafbc8c57186cb731d05c16010e35a445d77c782a5329ccfd2"

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("token") != token {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid token access"})
			c.Abort()
			return
		}
	}
}
