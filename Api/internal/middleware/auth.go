package middleware

import "github.com/gin-gonic/gin"

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 请求前

		c.Next()

	}
}
