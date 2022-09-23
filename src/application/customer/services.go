package customer

import (
	"github.com/AndrewVazzoler/dock-api-rest/src/application/customer/commands"
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/protocols"
	"github.com/AndrewVazzoler/dock-api-rest/src/shared"
)

type Queries struct{}

// Commands Contains all available command handlers of this app
type Commands struct {
	CreateCustomerHandler commands.CreateCustomerRequestHandler
	DeleteCustomerHandler commands.DeleteCustomerRequestHandler
}

type CustomerServices struct {
	Queries  Queries
	Commands Commands
}

func NewServices(
	ctx shared.Ctx, repo *protocols.AllRepositories) CustomerServices {
	return CustomerServices{
		Queries: Queries{},
		Commands: Commands{
			CreateCustomerHandler: commands.NewCreateRequestHandler(ctx, repo),
			DeleteCustomerHandler: commands.NewDeleteRequestHandler(ctx, repo),
		},
	}
}
