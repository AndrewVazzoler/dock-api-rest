package main

import (
	shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"
	"github.com/AndrewVazzoler/dock-api-rest/src/application"
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/config"
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/database"
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/server"
	"github.com/go-playground/validator/v10"
	validate_doc "github.com/paemuri/brdoc/v2"
)

func main() {
	cfg := config.New()
	ctx := shared.NewCtx()
	repositories := database.NewServices(ctx, cfg)
	app := application.NewServices(ctx, repositories)
	server := server.NewServer(ctx, app)
	server.ListenAndServe()
}

func ValidateCPF(fl validator.FieldLevel) bool {
	return validate_doc.IsCPF(fl.Field().String())
}
