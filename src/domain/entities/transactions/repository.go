package transactions

import "github.com/AndrewVazzoler/dock-api-rest/src/shared"

type TransactionRepository interface {
	Create(account *Transaction) (*Transaction, error)
	Update(account *Transaction) (*Transaction, error)
	FindByID(id string) (*Transaction, error)
	FindAll(pagination *shared.Pagination[[]*Transaction]) (*shared.Pagination[[]*Transaction], error)
	Delete(id string) error
}
