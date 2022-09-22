package application

import (
	shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"
	protocols "github.com/AndrewVazzoler/dock-api-rest/src/domain/_protocols"

	"github.com/AndrewVazzoler/dock-api-rest/src/application/account"
	"github.com/AndrewVazzoler/dock-api-rest/src/application/customer"
)

type AllApplications struct {
	CustomerServices customer.CustomerServices
	AccountServices  account.AccountServices
}

func NewServices(ctx shared.Ctx, repo *protocols.AllRepositories) AllApplications {
	return AllApplications{
		CustomerServices: customer.NewServices(ctx, repo),
		AccountServices:  account.NewServices(ctx, repo),
	}
}
