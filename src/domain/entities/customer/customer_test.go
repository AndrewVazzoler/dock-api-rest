package customer_test

import (
	"testing"

	"github.com/AndrewVazzoler/dock-api-rest/src/domain/entities/customer"
	"github.com/stretchr/testify/require"
)

func TestCustomerValid(t *testing.T) {
	_, err := customer.NewCustomer("Antonio Freitas", "285.263.810-08")
	require.Nil(t, err)
}

func TestCustomerDocumentInvalid(t *testing.T) {
	_, err := customer.NewCustomer("Antonio Freitas", "285.263.810-09")
	require.Error(t, err)
}

func TestCustomerNameInvalid(t *testing.T) {
	_, err := customer.NewCustomer("", "285.263.810-08")
	require.Error(t, err)
}
