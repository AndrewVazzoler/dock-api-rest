package customer

import shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"

type CustomerRepository interface {
	Create(customer *Customer) (*Customer, error)
	FindByID(id string) (*Customer, error)
	FindByDocument(document string) (*Customer, error)
	FindAll(pagination *shared.Pagination[[]*Customer]) (*shared.Pagination[[]*Customer], error)
	// Delete(id string) error
}
