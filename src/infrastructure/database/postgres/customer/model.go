package customer

import (
	"time"

	"github.com/AndrewVazzoler/dock-api-rest/src/domain/customer"
)

type Customer struct {
	ID        string    `gorm:"primary_key;column:id"`
	Name      string    `gorm:"column:name"`
	Document  string    `gorm:"column:document"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func ToModel(c *customer.Customer) *Customer {
	return &Customer{
		ID:        c.ID,
		Name:      c.Name,
		Document:  c.Document,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func ToEntity(c *Customer) *customer.Customer {
	return &customer.Customer{
		ID:        c.ID,
		Name:      c.Name,
		Document:  c.Document,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func ToEntityMany(f []*Customer) []*customer.Customer {
	items := make([]*customer.Customer, len(f))

	for i, item := range f {
		items[i] = ToEntity(item)
	}

	return items
}
