package commands

import (
	"strconv"

	shared "github.com/AndrewVazzoler/dock-api-rest/src/_shared"
	protocols "github.com/AndrewVazzoler/dock-api-rest/src/domain/_protocols"
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/customer"

	"github.com/AndrewVazzoler/dock-api-rest/src/domain/account"
)

type (
	OpenAccountRequest struct {
		Document string
	}
	OpenAccountRequestHandler interface {
		Handle(command OpenAccountRequest) (*account.Account, error)
	}
	openAccountRequestHandler struct {
		ctx          shared.Ctx
		repo         account.AccountRepository
		repoCustomer customer.CustomerRepository
	}
)

func NewAccountRequestHandler(ctx shared.Ctx, repo *protocols.AllRepositories) OpenAccountRequestHandler {
	return &openAccountRequestHandler{
		ctx:          ctx,
		repo:         repo.AccountRepository,
		repoCustomer: repo.CustomerRepository,
	}
}

func (h openAccountRequestHandler) Handle(req OpenAccountRequest) (*account.Account, error) {
	customer, err := h.repoCustomer.FindByDocument(req.Document)

	if err != nil {
		return nil, err
	}

	lastAccountNumber, err := h.repo.FindLastAccountNumber()

	if err != nil {
		lastAccountNumber = "00000"
	}

	i, err := strconv.ParseInt(lastAccountNumber, 10, 8)
	v := i + 1
	vq := strconv.Itoa(int(v))

	if err != nil {
		return nil, err
	}

	newAccount, err := account.NewAccount(customer, 0, "001", vq, true, false)

	if err != nil {
		return nil, err
	}

	result, err := h.repo.Create(newAccount)

	if err != nil {
		return nil, err
	}

	return result, nil
}
