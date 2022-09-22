package controllers

import (
	"errors"
	"fmt"

	shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"
	"github.com/AndrewVazzoler/dock-api-rest/src/application"
	"github.com/AndrewVazzoler/dock-api-rest/src/application/customer/commands"
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/server/gin/utils"
	"github.com/gin-gonic/gin"
)

type AccountRequest struct {
	Name     string `binding:"required" json:"name"`
	Document string `binding:"required,isCpf" json:"document"`
}

func CreateCustomer(ctx shared.Ctx, app application.AllApplications) gin.HandlerFunc {
	return func(c *gin.Context) {
		var model AccountRequest

		errBind := utils.BindBodyFromPostRequest(c, &model)

		if errBind != nil {
			ctx.Http.BadRequest(errors.New("validation"), errBind)
			return
		}

		result, err := app.CustomerServices.Commands.CreateCustomerHandler.Handle(
			commands.CreateCustomerRequest{
				Name:     model.Name,
				Document: model.Document,
			},
		)

		if err != nil {
			fmt.Print(err)
			ctx.Http.BadRequest(err, nil)
			return
		}

		ctx.Http.CreateOK(result)
	}
}
