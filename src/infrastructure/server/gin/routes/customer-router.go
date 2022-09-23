package routes

import (
	"github.com/AndrewVazzoler/dock-api-rest/src/application"
	controller "github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/server/gin/controllers/customer"
	"github.com/AndrewVazzoler/dock-api-rest/src/shared"

	"github.com/gin-gonic/gin"
)

func CustomerRouter(r *gin.Engine, ctx shared.Ctx, app application.AllApplications) {
	customers := r.Group("/customers")
	{
		customers.POST("", controller.CreateCustomer(ctx, app))
		customers.DELETE(":id", controller.DeleteCustomer(ctx, app))
	}
}
