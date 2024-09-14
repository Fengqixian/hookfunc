package middleware

import (
	"github.com/gin-gonic/gin"
	"hookfunc/pkg/log"
)

func UserRole(logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
