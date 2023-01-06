package domain

import (
	"database/sql"
)

type AccountRepository interface {
	Save(input *Account) error
	Update(input *Account) error
	Delete(id string) error
	GetOne(id string) (*AccountOutpuDTO, error)
	GetMany(accountType string) ([]*AccountOutpuDTO, error)
}

type AccountOutpuDTO struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Email       string       `json:"email"`
	CellPhone   string       `json:"cell_phone"`
	AccountType string       `json:"account_type"`
	CreditLimit float32      `json:"credit_limit"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}
