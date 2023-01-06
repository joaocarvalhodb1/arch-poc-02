package sql

import "database/sql"

type DBDriver interface {
	Close() error
	CreateStatement(sql string, params ...any) (*sql.Stmt, error)
	QueryExecute(sql string, params ...any) error
	QueryOpen(sql string, params ...any) (*sql.Rows, error)
}
