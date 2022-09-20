package main

import (
	shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"
	"github.com/AndrewVazzoler/dock-api-rest/src/application"
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/config"
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/database"
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/server"
)

func main() {
	cfg := config.New()
	ctx := shared.NewCtx()
	repositories := database.NewServices(ctx, cfg)
	app := application.NewServices(ctx, repositories)
	server := server.NewServer(ctx, app)
	server.ListenAndServe()
}
