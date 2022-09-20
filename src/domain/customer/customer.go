package customer

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	validate_doc "github.com/paemuri/brdoc/v2"
	uuid "github.com/satori/go.uuid"
)

type Customer struct {
	ID        string `validate:"uuid"`
	Name      string `validate:"required"`
	Document  string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCustomer(name string, document string) (*Customer, error) {
	customer := Customer{
		Name:     name,
		Document: document,
	}

	customer.prepare()

	err := customer.Validate()

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (customer *Customer) prepare() {
	customer.ID = uuid.NewV4().String()
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()
}

func (customer *Customer) Validate() error {
	validate := validator.New()
	err := validate.Struct(customer)

	isValid := validate_doc.IsCPF(customer.Document)

	if !isValid {
		return fmt.Errorf("it is not a valid document")
	}

	if err != nil {
		return err
	}

	return nil
}
