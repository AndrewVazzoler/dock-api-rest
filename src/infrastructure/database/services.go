package database

import (
	shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"
	protocols "github.com/AndrewVazzoler/dock-api-rest/src/domain/_protocols"
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/config"
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/database/postgres"
	customerRepository "github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/database/postgres/customer"
)

func NewServices(ctx shared.Ctx, cfg config.Config) *protocols.AllRepositories {
	var db = postgres.NewDbAndConnect(cfg)
	return &protocols.AllRepositories{
		CustomerRepository: customerRepository.NewCustomerRepository(ctx, db),
	}
}
