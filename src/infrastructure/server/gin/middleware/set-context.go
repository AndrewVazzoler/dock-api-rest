package middleware

import (
	shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"

	"github.com/gin-gonic/gin"
)

func SetContextMiddleware(ctx shared.Ctx) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx.Http.SetServerContext(c)
	}
}
