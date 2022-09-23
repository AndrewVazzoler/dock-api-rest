package transactions

import (
	"time"

	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	ID        string
	AccountID string
	Operation string
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTransaction(accountId string, amount float64, operation string) (*Transaction, error) {
	trans := Transaction{
		AccountID: accountId,
		Operation: operation,
		Amount:    amount,
	}
	trans.prepare()

	err := trans.Validate()

	if err != nil {
		return nil, err
	}

	return &trans, nil
}

func (t *Transaction) prepare() {
	t.ID = uuid.NewV4().String()
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}

func (t *Transaction) Validate() error {
	validate := validator.New()
	err := validate.Struct(t)

	if err != nil {
		return err
	}

	return nil
}
