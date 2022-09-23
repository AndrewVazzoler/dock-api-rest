package commands

import (
	"fmt"

	"github.com/AndrewVazzoler/dock-api-rest/src/domain/entities/customer"
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/protocols"
	"github.com/AndrewVazzoler/dock-api-rest/src/shared"
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
	_, err := h.repo.FindByDocument(req.Document)

	if err == nil {
		return nil, fmt.Errorf("exist a customer with document %s", req.Document)
	}

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
