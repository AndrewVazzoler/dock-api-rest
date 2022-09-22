package account

import (
	"fmt"
	"time"

	"github.com/AndrewVazzoler/dock-api-rest/src/domain/account"
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/customer"
)

type Account struct {
	ID            string            `gorm:"primary_key;column:id"`
	OwnerID       string            `gorm:"uniqueIndex;column:owner_id"`
	Owner         customer.Customer `gorm:"foreignkey:OwnerID"`
	Balance       int64             `gorm:"column:balance"`
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
		OwnerID:       c.Owner.ID,
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
	fmt.Println(c.Owner)
	return &account.Account{
		ID:            c.ID,
		Owner:         &c.Owner,
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
