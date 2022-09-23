package account

import (
	"time"

	"github.com/AndrewVazzoler/dock-api-rest/src/domain/entities/account"
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/entities/customer"
)

type Account struct {
	ID            string            `gorm:"primary_key;column:id"`
	OwnerID       string            `gorm:"uniqueIndex;column:owner_id"`
	Owner         customer.Customer `gorm:"foreignkey:OwnerID"`
	Balance       float64           `gorm:"column:balance"`
	AgencyNumber  string            `gorm:"column:agency_number"`
	AccountNumber string            `gorm:"column:account_number"`
	Active        bool              `gorm:"column:active"`
	Lock          bool              `gorm:"column:lock"`
	CreatedAt     time.Time         `gorm:"column:created_at"`
	UpdatedAt     time.Time         `gorm:"column:updated_at"`
}

func ToModel(c *account.Account) *Account {
	return &Account{
		ID:            c.ID,
		OwnerID:       c.OwnerID,
		Balance:       c.Balance,
		AgencyNumber:  c.AgencyNumber,
		AccountNumber: c.AccountNumber,
		Active:        c.Active,
		Lock:          c.Lock,
		CreatedAt:     c.CreatedAt,
		UpdatedAt:     c.UpdatedAt,
	}
}

func ToEntity(c *Account) *account.Account {
	return &account.Account{
		ID:            c.ID,
		OwnerID:       c.OwnerID,
		Balance:       c.Balance,
		AgencyNumber:  c.AgencyNumber,
		AccountNumber: c.AccountNumber,
		Active:        c.Active,
		Lock:          c.Lock,
		CreatedAt:     c.CreatedAt,
		UpdatedAt:     c.UpdatedAt,
	}
}

func ToEntityMany(f []*Account) []*account.Account {
	items := make([]*account.Account, len(f))

	for i, item := range f {
		items[i] = ToEntity(item)
	}

	return items
}
