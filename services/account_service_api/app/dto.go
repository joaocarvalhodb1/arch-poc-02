package app

import (
	"strings"

	"github.com/asaskevich/govalidator"
)

type CreditLimiteDTO struct {
	AccountType string  `json:"account_type" valid:"required, maxstringlength(10)"`
	CreditLimit float32 `json:"credit_limit" valid:"required, type(float32)"`
}

type AccountDTO struct {
	ID          string  `json:"id" valid:"-"`
	Name        string  `json:"name" valid:"required, maxstringlength(100)"`
	Email       string  `json:"email" valid:"email"`
	CellPhone   string  `json:"cell_phone" valid:"required, maxstringlength(15)"`
	AccountType string  `json:"account_type"  valid:"required, maxstringlength(10)"` // 0=LEAD, 1=NORMAL, 2=PREMIUM
	CreditLimit float32 `json:"credit_limit" valid:"required"`
}

func (a *AccountDTO) Validate() error {
	a.Name = strings.ToUpper(strings.TrimSpace(a.Name))
	a.Email = strings.ToLower(strings.TrimSpace(a.Email))
	_, err := govalidator.ValidateStruct(a)
	if err != nil {
		return err
	}
	return nil
}

func (c *CreditLimiteDTO) Validate() error {
	c.AccountType = strings.ToUpper(strings.TrimSpace(c.AccountType))
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}
	return nil
}
