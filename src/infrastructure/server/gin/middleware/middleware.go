package middleware

import (
	"github.com/AndrewVazzoler/dock-api-rest/src/shared"

	"github.com/gin-gonic/gin"
)

func UseMiddlewares(ctx shared.Ctx, r *gin.Engine) {
	r.Use(UseCORS())
	r.Use(DetectLanguage())
	r.Use(SetContextMiddleware(ctx))
}
