package infra

import (
	"strings"

	"github.com/joaocarvalhodb1/arch-poc/services/account_service/domain"
	"github.com/joaocarvalhodb1/arch-poc/shared/helpers"
	"github.com/joaocarvalhodb1/arch-poc/shared/sql"
)

type AccountRepoPostgres struct {
	driver sql.DBDriver
	log    *helpers.Loggers
}

func NewAccountRepoPostgres(driver sql.DBDriver, log *helpers.Loggers) domain.AccountRepository {
	repo := &AccountRepoPostgres{
		driver: driver,
		log:    log,
	}
	return repo
}

func (r *AccountRepoPostgres) Save(input *domain.Account) error {
	sql := `INSERT INTO accounts (id, name, email, cell_phone, account_type, credit_limit, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	err := r.driver.QueryExecute(sql, input.ID, input.Name, input.Email, input.CellPhone, input.AccountType, input.CreditLimit, input.CreatedAt)
	if err != nil {
		r.log.Error("Error on create account", err)
		return err
	}
	return nil
}

func (r *AccountRepoPostgres) Update(input *domain.Account) error {
	sql := `UPDATE accounts SET name = $1, email = $2, cell_phone = $3, credit_limit = $4, account_type = $5, updated_at = $6 WHERE id = $7`
	err := r.driver.QueryExecute(sql, input.Name, input.Email, input.CellPhone, input.CreditLimit, input.AccountType, input.UpdatedAt, input.ID)
	if err != nil {
		r.log.Error("Error on update account", err)
		return err
	}
	return nil
}

func (r *AccountRepoPostgres) Delete(id string) error {
	sql := `DELETE FROM accounts WHERE id = $1`
	err := r.driver.QueryExecute(sql, id)
	if err != nil {
		r.log.Error("Error on delete account", err)
		return err
	}
	return nil
}

func (r *AccountRepoPostgres) GetOne(id string) (*domain.AccountOutpuDTO, error) {
	sql := `SELECT id, name, email, cell_phone, account_type, credit_limit, created_at, updated_at from accounts WHERE id = $1`
	rows, err := r.driver.QueryOpen(sql, id)
	if err != nil {
		r.log.Error("Error on find a account", err)
		return nil, err
	}
	if rows.Err() != nil {
		r.log.Error("account not found", err)
		return nil, err
	}

	accounts := []*domain.AccountOutpuDTO{}
	for rows.Next() {
		var account domain.AccountOutpuDTO
		if err := rows.Scan(
			&account.Id,
			&account.Name,
			&account.Email,
			&account.CellPhone,
			&account.AccountType,
			&account.CreditLimit,
			&account.CreatedAt,
			&account.UpdatedAt,
		); err != nil {
			r.log.Error("Error on scan a account", err)
			return nil, err
		}
		accounts = append(accounts, &account)
	}
	defer rows.Close()
	return accounts[0], nil
}

func (r *AccountRepoPostgres) GetMany(accountType string) ([]*domain.AccountOutpuDTO, error) {
	if strings.TrimSpace(accountType) == "" {
		accountType = "%"
	}
	sql := `SELECT id, name, email, cell_phone, account_type, credit_limit, created_at, updated_at from accounts where account_type like $1`
	rows, err := r.driver.QueryOpen(sql, accountType)
	if err != nil {
		r.log.Error("Error on findMany accounts", err)
		return nil, err
	}
	accounts := []*domain.AccountOutpuDTO{}
	for rows.Next() {
		var account domain.AccountOutpuDTO
		if err := rows.Scan(
			&account.Id,
			&account.Name,
			&account.Email,
			&account.CellPhone,
			&account.AccountType,
			&account.CreditLimit,
			&account.CreatedAt,
			&account.UpdatedAt,
		); err != nil {
			r.log.Error("Error on scan a account", err)
			return nil, err
		}
		accounts = append(accounts, &account)
	}
	defer rows.Close()
	return accounts, nil
}
