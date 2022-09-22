package controllers

import (
	"errors"
	"time"

	shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"
	"github.com/AndrewVazzoler/dock-api-rest/src/application"
	"github.com/AndrewVazzoler/dock-api-rest/src/application/account/commands"
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/server/gin/utils"
	"github.com/gin-gonic/gin"
)

type OpenAccountRequest struct {
	Document string `binding:"required,isCpf" json:"document"`
}
type OpenAccountResponse struct {
	ID            string    `json:"id"`
	Balance       int64     `json:"balance"`
	AgencyNumber  string    `json:"agency_number"`
	AccountNumber string    `json:"account_number"`
	Active        bool      `json:"active"`
	Lock          bool      `json:"lock"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func OpenAccount(ctx shared.Ctx, app application.AllApplications) gin.HandlerFunc {
	return func(c *gin.Context) {
		var model OpenAccountRequest

		errBind := utils.BindBodyFromPostRequest(c, &model)

		if errBind != nil {
			ctx.Http.BadRequest(errors.New("validation"), errBind)
			return
		}

		result, err := app.AccountServices.Commands.OpenAccountHandler.Handle(
			commands.OpenAccountRequest{
				Document: model.Document,
			},
		)

		if err != nil {
			ctx.Http.BadRequest(err, nil)
			return
		}

		var response OpenAccountResponse
		ctx.ToolsDTO.Map(&response, result)
		ctx.Http.CreateOK(response)
	}
}
