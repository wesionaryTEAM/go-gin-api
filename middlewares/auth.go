package middlewares

import "github.com/gin-gonic/gin"

// BasicAuth is ...
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"Dinesh": "Silwal",
	})
}
