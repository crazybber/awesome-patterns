package main

import (
	"errors"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash      = 1
	DebitCard = 2
)

type CashPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%v paid using cash\n", amount)
}

type DebitCardPM struct{}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%v paid using DebitCard\n", amount)
}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d was not recognized"))
	}
}

func main() {
	paymentMethod, err := GetPaymentMethod(Cash)
	if err != nil {
		fmt.Sprintf("Error %s", err.Error())
	}
	spew.Dump(paymentMethod.Pay(32))
}
