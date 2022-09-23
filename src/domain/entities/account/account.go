package account

import (
	"time"

	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	ID            string  `validate:"uuid"`
	OwnerID       string  `validate:"required"`
	Balance       float64 `validate:"gte=0"`
	AgencyNumber  string  `validate:"required"`
	AccountNumber string  `validate:"required"`
	Active        bool
	Lock          bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewAccount(ownerID string, balance float64, agencyNumber string, accountNumber string, active bool, lock bool) (*Account, error) {
	account := Account{
		OwnerID:       ownerID,
		Balance:       balance,
		AgencyNumber:  agencyNumber,
		AccountNumber: accountNumber,
		Lock:          lock,
		Active:        active,
	}
	account.prepare()

	err := account.Validate()

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (account *Account) prepare() {
	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()
}

func (account *Account) Validate() error {
	validate := validator.New()
	err := validate.Struct(account)

	if err != nil {
		return err
	}

	return nil
}

func (account *Account) SetLock() {
	account.UpdatedAt = time.Now()
	account.Lock = true
}

func (account *Account) SetUnlock() {
	account.UpdatedAt = time.Now()
	account.Lock = false
}
func (account *Account) Desactive() {
	account.UpdatedAt = time.Now()
	account.Active = false
}
func (account *Account) IncrementBalance(amount float64) {
	account.Balance = account.Balance + amount
	account.UpdatedAt = time.Now()
}
