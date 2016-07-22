package bleh

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Bleh struct {
	db *sqlx.DB
}

func NewBleh(dbDriver string, dbURN string) (*Bleh, error) {
	var db *sqlx.DB
	db, err := sqlx.Connect(dbDriver, dbURN)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %s", err)
	}

	return &Bleh{
		db: db,
	}, nil
}

func (b *Bleh) Foo() (string, error) {
	return "this is foo.", nil
}

func (b *Bleh) GetAge(name string) (int, error) {
	var age int
	err := b.db.Get(&age, "select age from person where name = $1", name)
	if err != nil {
		return 0, err
	}

	return age, nil
}

func (b *Bleh) InitDB() error {
	_, err := b.db.Exec(schema)
	if err != nil {
		return fmt.Errorf("unable to execute schema: %s", err)
	}
	return nil
}

func (b *Bleh) DropDB() error {
	_, err := b.db.Exec(dropSchema)
	if err != nil {
		return fmt.Errorf("unable to drop schema: %s", err)
	}
	return nil
}
