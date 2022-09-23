package controllers

import (
	"errors"
	"time"

	"github.com/AndrewVazzoler/dock-api-rest/src/application"
	"github.com/AndrewVazzoler/dock-api-rest/src/application/account/commands"
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/server/gin/utils"
	"github.com/AndrewVazzoler/dock-api-rest/src/shared"
	"github.com/gin-gonic/gin"
)

type DepositRequest struct {
	Document string  `binding:"required,isCpf" json:"document"`
	Amount   float64 `binding:"required,gt=0" json:"amount"`
}

type DepositResponse struct {
	TransactionID string    `json:"transaction_id"`
	CreatedAt     time.Time `json:"created_at"`
}

func Deposit(ctx shared.Ctx, app application.AllApplications) gin.HandlerFunc {
	return func(c *gin.Context) {
		var model DepositRequest

		errBind := utils.BindBodyFromPostRequest(c, &model)

		if errBind != nil {
			ctx.Http.BadRequest(errors.New("validation"), errBind)
			return
		}

		result, err := app.AccountServices.Commands.DepositHandler.Handle(
			commands.DepositRequest{
				Document: model.Document,
				Amount:   model.Amount,
			},
		)

		if err != nil {
			ctx.Http.BadRequest(err, nil)
			return
		}

		var response DepositResponse
		ctx.ToolsDTO.Map(&response, result)
		ctx.Http.CreateOK(response)
	}
}
