package controllers

import (
	"errors"

	shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"
	"github.com/AndrewVazzoler/dock-api-rest/src/application"
	"github.com/AndrewVazzoler/dock-api-rest/src/application/account/commands"
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/server/gin/utils"
	"github.com/gin-gonic/gin"
)

type CloseAccountRequest struct {
	Document string `binding:"required,isCpf" json:"document"`
}

func CloseAccount(ctx shared.Ctx, app application.AllApplications) gin.HandlerFunc {
	return func(c *gin.Context) {
		var model CloseAccountRequest

		errBind := utils.BindBodyFromPostRequest(c, &model)

		if errBind != nil {
			ctx.Http.BadRequest(errors.New("validation"), errBind)
			return
		}

		_, err := app.AccountServices.Commands.CloseAccountHandler.Handle(
			commands.CloseAccountRequest{
				Document: model.Document,
			},
		)

		if err != nil {
			ctx.Http.BadRequest(err, nil)
			return
		}

		ctx.Http.DeleteOK()
	}
}
