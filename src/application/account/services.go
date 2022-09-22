package account

import (
	shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"
	"github.com/AndrewVazzoler/dock-api-rest/src/application/account/commands"
	protocols "github.com/AndrewVazzoler/dock-api-rest/src/domain/_protocols"
)

type Queries struct{}

// Commands Contains all available command handlers of this app
type Commands struct {
	OpenAccountHandler  commands.OpenAccountRequestHandler
	CloseAccountHandler commands.CloseAccountRequestHandler
}

type AccountServices struct {
	Queries  Queries
	Commands Commands
}

func NewServices(
	ctx shared.Ctx, repo *protocols.AllRepositories) AccountServices {
	return AccountServices{
		Queries: Queries{},
		Commands: Commands{
			OpenAccountHandler:  commands.NewAccountRequestHandler(ctx, repo),
			CloseAccountHandler: commands.NewCloseAccountRequestHandler(ctx, repo),
		},
	}
}
