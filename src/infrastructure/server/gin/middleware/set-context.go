package middleware

import (
	"github.com/AndrewVazzoler/dock-api-rest/src/shared"

	"github.com/gin-gonic/gin"
)

func SetContextMiddleware(ctx shared.Ctx) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx.Http.SetServerContext(c)
	}
}
