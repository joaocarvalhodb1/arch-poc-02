package domain

import (
	"fmt"

	"github.com/joaocarvalhodb1/arch-poc/shared/helpers"
)

type ApplyCreditLimit struct {
	accountRepository AccountRepository
	log               *helpers.Loggers
}

func NewApplyCreditLimit(accountRepository AccountRepository) *ApplyCreditLimit {
	creditLimit := &ApplyCreditLimit{accountRepository: accountRepository}
	return creditLimit
}

func (c *ApplyCreditLimit) Apply(accountType string, value float32) error {
	accounts, err := c.accountRepository.GetMany(accountType)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	for _, outputAccount := range accounts {
		account := convertOutputToAccount(outputAccount)
		err := account.ApplyCreditLimit(value)
		if err != nil {
			c.log.Error("Error on apply Credit limit", err)
			return fmt.Errorf(err.Error())
		}
		err = c.accountRepository.Update(account)
		if err != nil {
			c.log.Error("Error on apply Credit limit", err)
			return fmt.Errorf(err.Error())
		}
	}
	return nil
}

func convertOutputToAccount(input *AccountOutpuDTO) *Account {
	return &Account{
		ID:          input.Id,
		Name:        input.Name,
		Email:       input.Email,
		CellPhone:   input.CellPhone,
		AccountType: input.AccountType,
		CreditLimit: input.CreditLimit,
		UpdatedAt:   input.UpdatedAt,
	}
}
