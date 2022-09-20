package commands

import (
	shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"
	protocols "github.com/AndrewVazzoler/dock-api-rest/src/domain/_protocols"

	"github.com/AndrewVazzoler/dock-api-rest/src/domain/customer"
)

type (
	CreateCustomerRequest struct {
		Name     string
		Document string
	}
	CreateCustomerRequestHandler interface {
		Handle(command CreateCustomerRequest) (*customer.Customer, error)
	}
	createCustomerRequestHandler struct {
		ctx  shared.Ctx
		repo customer.CustomerRepository
	}
)

func NewCreateRequestHandler(ctx shared.Ctx, repo *protocols.AllRepositories) CreateCustomerRequestHandler {
	return &createCustomerRequestHandler{
		ctx:  ctx,
		repo: repo.CustomerRepository,
	}
}

func (h createCustomerRequestHandler) Handle(req CreateCustomerRequest) (*customer.Customer, error) {
	newCustomer, err := customer.NewCustomer(req.Name, req.Document)

	if err != nil {
		return nil, err
	}

	result, err := h.repo.Create(newCustomer)

	if err != nil {
		return nil, err
	}
	return result, nil
}
