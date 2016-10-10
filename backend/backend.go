package backend

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Backend struct {
	db *sqlx.DB
}

func New(dbURN string) (*Backend, error) {
	var db *sqlx.DB
	db, err := sqlx.Connect("postgres", dbURN)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %s", err)
	}

	return &Backend{
		db: db,
	}, nil
}

func (b *Backend) Foo() (string, error) {
	return "this is foo.", nil
}

func (b *Backend) InitDB() error {
	_, err := b.db.Exec(schema)
	if err != nil {
		return fmt.Errorf("unable to execute schema: %s", err)
	}
	return nil
}

func (b *Backend) DropDB() error {
	_, err := b.db.Exec(dropSchema)
	if err != nil {
		return fmt.Errorf("unable to drop schema: %s", err)
	}
	return nil
}
