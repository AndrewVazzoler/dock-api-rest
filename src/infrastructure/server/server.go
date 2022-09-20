package server

import (
	shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"
	"github.com/AndrewVazzoler/dock-api-rest/src/application"
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/server/gin/middleware"
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/server/gin/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(ctx shared.Ctx, appServices application.AllApplications) *Server {
	engine := gin.New()
	httpServer := &Server{engine: engine}
	middleware.UseMiddlewares(ctx, httpServer.engine)
	routes.StartRoutes(httpServer.engine, ctx, appServices)
	return httpServer
}

func (httpServer *Server) ListenAndServe(addr ...string) {
	if len(addr) > 0 {
		httpServer.engine.Run(addr[0])
	} else {
		httpServer.engine.Run()
	}
}
