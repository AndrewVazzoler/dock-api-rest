package account

import (
	"github.com/AndrewVazzoler/dock-api-rest/src/application/account/commands"
	"github.com/AndrewVazzoler/dock-api-rest/src/application/account/queries"
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/protocols"
	"github.com/AndrewVazzoler/dock-api-rest/src/shared"
)

type Queries struct {
	GetBalanceAccountHandler queries.GetBalanceAccountRequestHandler
}

// Commands Contains all available command handlers of this app
type Commands struct {
	OpenAccountHandler  commands.OpenAccountRequestHandler
	CloseAccountHandler commands.CloseAccountRequestHandler
	DepositHandler      commands.DepositRequestHandler
}

type AccountServices struct {
	Queries  Queries
	Commands Commands
}

func NewServices(
	ctx shared.Ctx, repo *protocols.AllRepositories) AccountServices {
	return AccountServices{
		Queries: Queries{
			GetBalanceAccountHandler: queries.NewGetBalanceAccountRequestHandler(ctx, repo),
		},
		Commands: Commands{
			OpenAccountHandler:  commands.NewAccountRequestHandler(ctx, repo),
			CloseAccountHandler: commands.NewCloseAccountRequestHandler(ctx, repo),
			DepositHandler:      commands.NewDepositRequestHandler(ctx, repo),
		},
	}
}
