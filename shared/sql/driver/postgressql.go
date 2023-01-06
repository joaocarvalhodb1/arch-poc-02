package driver

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/joaocarvalhodb1/arch-poc/shared/helpers"
	_ "github.com/lib/pq"
)

var lock = &sync.Mutex{}

const DbTimeout = time.Second * 3

type PostgresSQLDriver struct {
	log *helpers.Loggers
	Db  *sql.DB
	ctx context.Context
}

func NewPostgresSQLDriver(dsn string, log *helpers.Loggers, ctx context.Context) (*PostgresSQLDriver, error) {
	lock.Lock()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
		return nil, err
	}
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error connecting to database Ping: " + err.Error())
		return nil, err
	}
	connection := &PostgresSQLDriver{
		log: log,
		Db:  db,
		ctx: ctx,
	}
	defer lock.Unlock()
	return connection, nil
}

func (driver *PostgresSQLDriver) Close() error {
	err := driver.Db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (driver *PostgresSQLDriver) CreateStatement(sql string, params ...any) (*sql.Stmt, error) {
	lock.Lock()
	defer lock.Unlock()
	statement, err := driver.Db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	return statement, nil
}

func (driver *PostgresSQLDriver) QueryExecute(sql string, params ...any) error {
	row := driver.Db.QueryRowContext(driver.ctx, sql, params...)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (driver *PostgresSQLDriver) QueryOpen(sql string, params ...any) (*sql.Rows, error) {
	row, err := driver.Db.QueryContext(driver.ctx, sql, params...)
	if err != nil {
		return nil, err
	}
	return row, nil
}
