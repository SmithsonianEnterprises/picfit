package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/picfit/config"
)

func AddHeader(headers []config.Header) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, headerPair := range headers {
			c.Header(headerPair.Name, headerPair.Value)
		}
		c.Next()
	}
}
