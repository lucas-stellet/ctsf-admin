package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	Base   `valid:"required"`
	Type   string  `json:"name" valid:"notnull"`
	Amount float64 `json:"Amount" valid:"notnull"`
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)

	if err != nil {
		return err
	}

	return nil
}

func (t Transaction) NewTransaction(transactionType string, amount float64) (*Transaction, error) {
	transaction := Transaction{
		Type:   transactionType,
		Amount: amount,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
