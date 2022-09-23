package account_test

import (
	"fmt"
	"testing"

	"github.com/AndrewVazzoler/dock-api-rest/src/domain/entities/account"
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/entities/customer"

	"github.com/stretchr/testify/require"
)

func TestAccountValid(t *testing.T) {
	customer, err := customer.NewCustomer("Antonio Freitas", "285.263.810-08")
	require.Nil(t, err)

	_, err = account.NewAccount(customer.ID, 0, "001", "4004", true, false)
	fmt.Print(err)
	require.Nil(t, err)
}

func TestAccountInvalid(t *testing.T) {
	customer, err := customer.NewCustomer("Antonio Freitas", "285.263.810-08")
	require.Nil(t, err)

	_, err = account.NewAccount(customer.ID, -2, "001", "4004", true, false)
	require.Error(t, err)
}
