package account

import "github.com/AndrewVazzoler/dock-api-rest/src/shared"

type AccountRepository interface {
	Create(account *Account) (*Account, error)
	Update(account *Account) (*Account, error)
	FindByID(id string) (*Account, error)
	FindByOwner(id string) (*Account, error)
	FindLastAccountNumber() (string, error)
	FindAll(pagination *shared.Pagination[[]*Account]) (*shared.Pagination[[]*Account], error)
	Delete(id string) error
}
