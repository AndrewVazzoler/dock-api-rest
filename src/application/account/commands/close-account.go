package commands

import (
	shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"
	protocols "github.com/AndrewVazzoler/dock-api-rest/src/domain/_protocols"
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/customer"

	"github.com/AndrewVazzoler/dock-api-rest/src/domain/account"
)

type (
	CloseAccountRequest struct {
		Document string
	}
	CloseAccountRequestHandler interface {
		Handle(command CloseAccountRequest) (*account.Account, error)
	}
	closeAccountRequestHandler struct {
		ctx          shared.Ctx
		repo         account.AccountRepository
		repoCustomer customer.CustomerRepository
	}
)

func NewCloseAccountRequestHandler(ctx shared.Ctx, repo *protocols.AllRepositories) CloseAccountRequestHandler {
	return &closeAccountRequestHandler{
		ctx:          ctx,
		repo:         repo.AccountRepository,
		repoCustomer: repo.CustomerRepository,
	}
}

func (h closeAccountRequestHandler) Handle(req CloseAccountRequest) (*account.Account, error) {
	customer, err := h.repoCustomer.FindByDocument(req.Document)

	if err != nil {
		return nil, err
	}

	result, err := h.repo.FindByOwner(customer.ID)

	if err != nil {
		return nil, err
	}

	result.Desactive()

	result, err = h.repo.Update(result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
