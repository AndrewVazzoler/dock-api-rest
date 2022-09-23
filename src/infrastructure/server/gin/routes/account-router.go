package routes

import (
	"github.com/AndrewVazzoler/dock-api-rest/src/application"
	controller "github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/server/gin/controllers/account"
	"github.com/AndrewVazzoler/dock-api-rest/src/shared"

	"github.com/gin-gonic/gin"
)

func AccountRouter(r *gin.Engine, ctx shared.Ctx, app application.AllApplications) {
	accounts := r.Group("/accounts")
	{
		accounts.POST("open", controller.OpenAccount(ctx, app))
		accounts.POST("close", controller.CloseAccount(ctx, app))
		accounts.POST("deposit", controller.Deposit(ctx, app))
		accounts.POST("balance", controller.GetBalanceAccount(ctx, app))
	}
}
