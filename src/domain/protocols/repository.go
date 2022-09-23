package protocols

import (
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/entities/account"
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/entities/customer"
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/entities/transactions"
)

type AllRepositories struct {
	CustomerRepository    customer.CustomerRepository
	AccountRepository     account.AccountRepository
	TransactionRepository transactions.TransactionRepository
}
