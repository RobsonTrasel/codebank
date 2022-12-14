package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CreditCard struct {
	ID              string
	Name            string
	Number          string
	ExpirationMonth int32
	ExpirationYear  int32
	CCV             int32
	Balance         float64
	Limit           float64
	CreatedAt       time.Time
}

func newCreditCard() *CreditCard {
	c := &CreditCard{}
	c.ID = uuid.NewV1().String()
	c.CreatedAt = time.Now()
	return c
}
