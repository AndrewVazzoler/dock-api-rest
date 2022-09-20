package account_test

import (
	"fmt"
	"testing"

	"github.com/AndrewVazzoler/dock-api-rest/src/domain/account"
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/customer"

	"github.com/stretchr/testify/require"
)

func TestAccountValid(t *testing.T) {
	customer, err := customer.NewCustomer("Antonio Freitas", "285.263.810-08")
	require.Nil(t, err)

	_, err = account.NewAccount(customer, 0, "001", "4004")
	fmt.Print(err)
	require.Nil(t, err)
}

func TestAccountInvalid(t *testing.T) {
	customer, err := customer.NewCustomer("Antonio Freitas", "285.263.810-08")
	require.Nil(t, err)

	_, err = account.NewAccount(customer, -2, "001", "4004")
	require.Error(t, err)
}
