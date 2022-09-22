package commands

import (
	"fmt"

	shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"
	protocols "github.com/AndrewVazzoler/dock-api-rest/src/domain/_protocols"

	"github.com/AndrewVazzoler/dock-api-rest/src/domain/customer"
)

type (
	DeleteCustomerRequest struct {
		Id string
	}
	DeleteCustomerRequestHandler interface {
		Handle(command DeleteCustomerRequest) error
	}
	deleteCustomerRequestHandler struct {
		ctx  shared.Ctx
		repo customer.CustomerRepository
	}
)

func NewDeleteRequestHandler(ctx shared.Ctx, repo *protocols.AllRepositories) DeleteCustomerRequestHandler {
	return &deleteCustomerRequestHandler{
		ctx:  ctx,
		repo: repo.CustomerRepository,
	}
}

func (h deleteCustomerRequestHandler) Handle(req DeleteCustomerRequest) error {
	_, err := h.repo.FindByID(req.Id)

	if err != nil {
		return fmt.Errorf("not exist a customer with id %s", req.Id)
	}

	err = h.repo.Delete(req.Id)

	if err != nil {
		return err
	}

	return nil
}
