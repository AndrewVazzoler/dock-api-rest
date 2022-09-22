package protocols

import (
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/account"
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/customer"
)

type AllRepositories struct {
	CustomerRepository customer.CustomerRepository
	AccountRepository  account.AccountRepository
}
