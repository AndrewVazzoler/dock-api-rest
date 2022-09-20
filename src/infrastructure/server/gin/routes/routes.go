package routes

import (
	shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"
	"github.com/AndrewVazzoler/dock-api-rest/src/application"

	"github.com/gin-gonic/gin"
)

func StartRoutes(e *gin.Engine, ctx shared.Ctx, app application.AllApplications) {
	HealthCheck(e)
}
