package database

import (
	"github.com/dropbox/godropbox/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/seesawlabs/crawlfish/shared/config"
)

type database struct {
	Dbx *sqlx.DB
}

func (d *database) Sqlx() *sqlx.DB {
	return d.Dbx
}

type IDatabase interface {
	Sqlx() *sqlx.DB
}

func NewDatabase(c config.Database) (IDatabase, error) {
	db := &database{}

	var err error
	if db.Dbx, err = sqlx.Connect("mysql", dsn(c)); err != nil {
		return nil, errors.Wrap(err, "")
	}

	if err := db.Dbx.Ping(); err != nil {
		return nil, errors.Wrap(err, "")
	}

	return db, nil
}

// dsn prepare dsn connection string
// Example: root:@tcp(localhost:3306)/test
func dsn(c config.Database) string {
	return c.Username +
		":" +
		c.Password +
		"@tcp(" +
		c.Hostname +
		":" +
		c.Port +
		")/" +
		c.Name + c.Parameter
}
