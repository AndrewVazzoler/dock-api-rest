package database

import (
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/protocols"
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/config"
	accountRepository "github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/database/postgres/account"
	customerRepository "github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/database/postgres/customer"
	transactionRepository "github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/database/postgres/transaction"
	"github.com/AndrewVazzoler/dock-api-rest/src/shared"
)

func NewServices(ctx shared.Ctx, cfg config.Config) *protocols.AllRepositories {
	var db = NewDbAndConnect(cfg)
	return &protocols.AllRepositories{
		CustomerRepository:    customerRepository.NewCustomerRepository(ctx, db),
		AccountRepository:     accountRepository.NewAccountRepository(ctx, db),
		TransactionRepository: transactionRepository.NewTransactionRepository(ctx, db),
	}
}
