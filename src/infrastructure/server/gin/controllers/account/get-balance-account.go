package controllers

import (
	"errors"

	"github.com/AndrewVazzoler/dock-api-rest/src/application"
	"github.com/AndrewVazzoler/dock-api-rest/src/application/account/queries"
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/server/gin/utils"
	"github.com/AndrewVazzoler/dock-api-rest/src/shared"
	"github.com/gin-gonic/gin"
)

type GetBalanceAccountRequest struct {
	Document string `binding:"required,isCpf" json:"document"`
}

type GetBalanceAccountResponse struct {
	Balance float64 `json:"balance"`
}

func GetBalanceAccount(ctx shared.Ctx, app application.AllApplications) gin.HandlerFunc {
	return func(c *gin.Context) {
		var model GetBalanceAccountRequest

		errBind := utils.BindBodyFromPostRequest(c, &model)

		if errBind != nil {
			ctx.Http.BadRequest(errors.New("validation"), errBind)
			return
		}

		result, err := app.AccountServices.Queries.GetBalanceAccountHandler.Handle(
			queries.GetBalanceAccountRequest{
				Document: model.Document,
			},
		)

		if err != nil {
			ctx.Http.BadRequest(err, nil)
			return
		}

		var response GetBalanceAccountResponse
		ctx.ToolsDTO.Map(&response, result)
		ctx.Http.OK(response)
	}
}
