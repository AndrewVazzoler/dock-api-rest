package middleware

import (
	shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"

	"github.com/gin-gonic/gin"
)

func UseMiddlewares(ctx shared.Ctx, r *gin.Engine) {
	r.Use(UseCORS())
	r.Use(DetectLanguage())
	r.Use(SetContextMiddleware(ctx))
}
