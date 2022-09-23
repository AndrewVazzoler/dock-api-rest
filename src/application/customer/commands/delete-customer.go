package commands

import (
	"fmt"

	"github.com/AndrewVazzoler/dock-api-rest/src/domain/entities/customer"
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/protocols"
	"github.com/AndrewVazzoler/dock-api-rest/src/shared"
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
