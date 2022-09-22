package routes

import (
	shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"
	"github.com/AndrewVazzoler/dock-api-rest/src/application"
	controller "github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/server/gin/controllers/account"

	"github.com/gin-gonic/gin"
)

func AccountRouter(r *gin.Engine, ctx shared.Ctx, app application.AllApplications) {
	accounts := r.Group("/accounts")
	{
		accounts.POST("open", controller.OpenAccount(ctx, app))
		accounts.POST("close", controller.CloseAccount(ctx, app))
	}
}
