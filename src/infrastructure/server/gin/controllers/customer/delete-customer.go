package controllers

import (
	"fmt"

	shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"
	"github.com/AndrewVazzoler/dock-api-rest/src/application"
	"github.com/AndrewVazzoler/dock-api-rest/src/application/customer/commands"
	"github.com/gin-gonic/gin"
)

func DeleteCustomer(ctx shared.Ctx, app application.AllApplications) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := app.CustomerServices.Commands.DeleteCustomerHandler.Handle(
			commands.DeleteCustomerRequest{
				Id: id,
			},
		)

		if err != nil {
			fmt.Print(err)
			ctx.Http.ResponseError(err)
			return
		}
		ctx.Http.DeleteOK()
	}
}
