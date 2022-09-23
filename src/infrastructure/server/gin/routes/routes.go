package routes

import (
	"github.com/AndrewVazzoler/dock-api-rest/src/application"
	"github.com/AndrewVazzoler/dock-api-rest/src/shared"

	"github.com/gin-gonic/gin"
)

func StartRoutes(e *gin.Engine, ctx shared.Ctx, app application.AllApplications) {
	HealthCheck(e)
	CustomerRouter(e, ctx, app)
	AccountRouter(e, ctx, app)
}
