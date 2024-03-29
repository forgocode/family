package middleware

import (
	"github.com/forgocode/family/internal/webservice/router/manager"
	"github.com/gin-gonic/gin"
)

func RouterManager() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if manager.ISRouterPass(ctx.FullPath()) {
			ctx.Next()
		}
		ctx.Abort()

	}
}
