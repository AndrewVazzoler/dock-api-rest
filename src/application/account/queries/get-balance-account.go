package queries

import (
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/entities/account"
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/entities/customer"
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/protocols"
	"github.com/AndrewVazzoler/dock-api-rest/src/shared"
)

type (
	GetBalanceAccountRequest struct {
		Document string
	}
	GetBalanceAccountResponse struct {
		Balance float64
	}
	GetBalanceAccountRequestHandler interface {
		Handle(command GetBalanceAccountRequest) (*GetBalanceAccountResponse, error)
	}
	getBalanceAccountRequestHandler struct {
		ctx          shared.Ctx
		repo         account.AccountRepository
		repoCustomer customer.CustomerRepository
	}
)

func NewGetBalanceAccountRequestHandler(ctx shared.Ctx, repo *protocols.AllRepositories) GetBalanceAccountRequestHandler {
	return &getBalanceAccountRequestHandler{
		ctx:          ctx,
		repo:         repo.AccountRepository,
		repoCustomer: repo.CustomerRepository,
	}
}

func (h getBalanceAccountRequestHandler) Handle(req GetBalanceAccountRequest) (*GetBalanceAccountResponse, error) {
	customer, err := h.repoCustomer.FindByDocument(req.Document)

	if err != nil {
		return nil, err
	}

	result, err := h.repo.FindByOwner(customer.ID)

	if err != nil {
		return nil, err
	}

	return &GetBalanceAccountResponse{
		Balance: result.Balance,
	}, nil
}
