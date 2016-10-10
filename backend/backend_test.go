package backend

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/bitterfly/kuho/spiderdata"
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

func TestBackend_Fill(t *testing.T) {
	backend := openBackend(t)
	defer closeBackend(t, backend)

	data := &spiderdata.Request{}
	unmarshalJSON(t, data, "sample_requests/simple_request.json")
	err := backend.Fill(data)
	if err != nil {
		t.Errorf("Could not fill database %s", err)
	}
	pause()
}

func pause() {
	sigs := make(chan os.Signal, 1)
	fmt.Printf("Press CTRL-C to continue...\n")
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
