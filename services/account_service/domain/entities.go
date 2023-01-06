package domain

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

const (
	LEAD    = "LEAD"
	NORMAL  = "NORMAL"
	PREMIUM = "PREMIUM"
)

type Account struct {
	ID          string       `json:"id" valid:"-"`
	Name        string       `json:"name" valid:"required, maxstringlength(100)"`
	Email       string       `json:"email" valid:"email"`
	CellPhone   string       `json:"cell_phone" valid:"required, maxstringlength(100)"`
	AccountType string       `json:"account_type" valid:"required, maxstringlength(10)"` // 0=LEAD, 1=NORMAL, 2=PREMIUM
	CreditLimit float32      `json:"credit_limit" valid:"type(float32)"`
	CreatedAt   sql.NullTime `json:"created_at" valid:"-"`
	UpdatedAt   sql.NullTime `json:"updated_at" valid:"-"`
}

func NewAccount(name, email, cell_phone string, creditLimit float32) (*Account, error) {
	account := &Account{
		ID:          uuid.New().String(),
		Name:        name,
		Email:       email,
		CellPhone:   cell_phone,
		AccountType: "LEAD",
		CreditLimit: creditLimit,
	}
	err := account.Validate()
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (c *Account) Validate() error {
	c.Name = strings.ToUpper(strings.TrimSpace(c.Name))
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Account) ApplyCreditLimit(value float32) error {
	if value < 0 {
		return fmt.Errorf("Invalid credit limit value: %f", value)
	}
	c.CreditLimit = value
	return nil
}
