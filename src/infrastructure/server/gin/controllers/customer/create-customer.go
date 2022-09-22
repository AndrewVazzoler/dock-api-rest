package controllers

import (
	"errors"
	"fmt"
	"time"

	shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"
	"github.com/AndrewVazzoler/dock-api-rest/src/application"
	"github.com/AndrewVazzoler/dock-api-rest/src/application/customer/commands"
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/server/gin/utils"
	"github.com/gin-gonic/gin"
)

type CreateCustomerRequest struct {
	Name     string `binding:"required" json:"name"`
	Document string `binding:"required,isCpf" json:"document"`
}

type CreateAccountResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Document  string    `json:"document"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateCustomer(ctx shared.Ctx, app application.AllApplications) gin.HandlerFunc {
	return func(c *gin.Context) {
		var model CreateCustomerRequest

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

		var response CreateAccountResponse
		ctx.ToolsDTO.Map(&response, result)
		ctx.Http.CreateOK(response)
	}
}
