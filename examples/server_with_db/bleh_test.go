package bleh

import (
	"flag"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

var dbURN string

func init() {
	var (
		dbName string
		dbUser string
	)

	flag.StringVar(&dbName, "db.name", "", "name of database to connect to")
	flag.StringVar(&dbUser, "db.user", "", "username to connect to database with")

	flag.Parse()

	dbURN = fmt.Sprintf(
		"user=%s dbname=%s sslmode=disable", dbUser, dbName,
	)
}

func TestBleh_Foo(t *testing.T) {
	assert := assert.New(t)

	bleh, err := NewBleh("postgres", dbURN)
	if err != nil {
		t.Fatalf("cannot create bleh: %s", err)
	}

	foo, err := bleh.Foo()
	if err != nil {
		t.Error(err)
	}

	assert.Equal("this is foo.", foo)
}

func TestBleh_GetAge(t *testing.T) {
	assert := assert.New(t)

	bleh, err := NewBleh("postgres", dbURN)
	if err != nil {
		t.Fatalf("cannot create bleh: %s", err)
	}

	bleh.InitDB()
	defer bleh.DropDB()

	db := sqlx.MustConnect("postgres", dbURN)
	db.MustExec(`insert into person(name, age) values('pesho', 42)`)

	age, err := bleh.GetAge("pesho")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(42, age)
}
