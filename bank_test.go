package midtrans_test

import (
	"testing"

	midtrans "github.com/ianeinser/go-midtrans"

	"github.com/cheekybits/is"
)

func TestBank(t *testing.T) {
	is := is.New(t)
	is.Equal("bni", midtrans.BankBni)
	is.Equal("mandiri", midtrans.BankMandiri)
	is.Equal("cimb", midtrans.BankCimb)
	is.Equal("bca", midtrans.BankBca)
	is.Equal("bri", midtrans.BankBri)
	is.Equal("maybank", midtrans.BankMaybank)
	is.Equal("permata", midtrans.BankPermata)
}
