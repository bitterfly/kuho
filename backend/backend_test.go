package backend

import (
	"flag"
	"fmt"
	"testing"
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

func openBackend(t *testing.T) *Backend {
	backend, err := New(dbURN)
	if err != nil {
		t.Fatalf("cannot create backend: %s", err)
	}

	err = backend.InitDB()
	if err != nil {
		t.Errorf("cannot create database: %s", err)
	}

	return backend
}

func closeBackend(t *testing.T, backend *Backend) {
	err := backend.DropDB()
	if err != nil {
		t.Errorf("cannot drop database: %s", err)
	}
}

func TestBackend_Info(t *testing.T) {
	backend := openBackend(t)
	defer closeBackend(t, backend)
}
