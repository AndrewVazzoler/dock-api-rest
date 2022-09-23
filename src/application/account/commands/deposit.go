package commands

import (
	"fmt"
	"time"

	"github.com/AndrewVazzoler/dock-api-rest/src/domain/entities/account"
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/entities/customer"
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/entities/transactions"
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/protocols"
	"github.com/AndrewVazzoler/dock-api-rest/src/shared"
)

type (
	DepositRequest struct {
		Document string
		Amount   float64
	}
	DepositResponse struct {
		TransactionID string
		CreatedAt     time.Time
	}
	DepositRequestHandler interface {
		Handle(command DepositRequest) (*DepositResponse, error)
	}
	depositRequestHandler struct {
		ctx              shared.Ctx
		repo             account.AccountRepository
		repoCustomer     customer.CustomerRepository
		repoTransactions transactions.TransactionRepository
	}
)

func NewDepositRequestHandler(ctx shared.Ctx, repo *protocols.AllRepositories) DepositRequestHandler {
	return &depositRequestHandler{
		ctx:              ctx,
		repo:             repo.AccountRepository,
		repoCustomer:     repo.CustomerRepository,
		repoTransactions: repo.TransactionRepository,
	}
}

func (h depositRequestHandler) Handle(req DepositRequest) (*DepositResponse, error) {
	customer, err := h.repoCustomer.FindByDocument(req.Document)

	if err != nil {
		return nil, err
	}

	result, err := h.repo.FindByOwner(customer.ID)

	if err != nil {
		return nil, err
	}

	result.IncrementBalance(req.Amount)
	fmt.Println(result)
	result, err = h.repo.Update(result)

	if err != nil {
		return nil, err
	}

	trans, err := transactions.NewTransaction(result.ID, req.Amount, "deposit")

	if err != nil {
		return nil, err
	}

	_, err = h.repoTransactions.Create(trans)

	if err != nil {
		return nil, err
	}

	return &DepositResponse{
		TransactionID: trans.ID,
		CreatedAt:     trans.CreatedAt,
	}, nil
}
