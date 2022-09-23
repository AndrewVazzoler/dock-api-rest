package transaction

import (
	"time"

	"github.com/AndrewVazzoler/dock-api-rest/src/domain/entities/transactions"
)

type Transaction struct {
	ID        string    `gorm:"primary_key;column:id"`
	AccountID string    `gorm:"column:account_id"`
	Amount    float64   `gorm:"column:amount"`
	Operation string    `gorm:"column:operation"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func ToModel(c *transactions.Transaction) *Transaction {
	return &Transaction{
		ID:        c.ID,
		AccountID: c.AccountID,
		Amount:    c.Amount,
		Operation: c.Operation,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func ToEntity(c *Transaction) *transactions.Transaction {
	return &transactions.Transaction{
		ID:        c.ID,
		AccountID: c.AccountID,
		Amount:    c.Amount,
		Operation: c.Operation,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func ToEntityMany(f []*Transaction) []*transactions.Transaction {
	items := make([]*transactions.Transaction, len(f))

	for i, item := range f {
		items[i] = ToEntity(item)
	}

	return items
}
