package tools

import (
	"github.com/gin-gonic/gin"
	"time"
)

func RateLimit() gin.HandlerFunc {
	return func(context *gin.Context) {
		limiter := time.Tick(500 * time.Millisecond)
		<-limiter
		context.Next()
	}
}
