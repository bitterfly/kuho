package bleh

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestBlehServer_Foo(t *testing.T) {
	assert := assert.New(t)

	bleh, err := NewBleh("postgres", dbURN)
	if err != nil {
		t.Fatalf("cannot create bleh: %s", err)
	}

	server := NewBlehServer(bleh)

	response := makeOkRequest(t, server, "GET", "/foo", nil)

	assert.Equal("this is foo.", string(response))
}

func TestBlehServer_GetAge(t *testing.T) {
	assert := assert.New(t)

	bleh, err := NewBleh("postgres", dbURN)
	if err != nil {
		t.Fatalf("cannot create bleh: %s", err)
	}

	bleh.InitDB()
	defer bleh.DropDB()

	db := sqlx.MustConnect("postgres", dbURN)
	db.MustExec(`insert into person(name, age) values('pesho', 42)`)

	server := NewBlehServer(bleh)

	response := makeOkRequest(t, server, "GET", "/get_age?person=pesho", nil)
	assert.Equal("pesho's age is 42", string(response))
}

func makeRequest(
	t *testing.T,
	handler http.Handler,
	method string, url string, body []byte,
) *httptest.ResponseRecorder {
	writer := httptest.NewRecorder()

	request, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		t.Fatalf("cannot create http request: %s", err)
	}

	handler.ServeHTTP(writer, request)

	return writer
}

func makeOkRequest(
	t *testing.T,
	handler http.Handler,
	method string, url string, body []byte,
) []byte {
	req := makeRequest(t, handler, method, url, body)
	if req.Code != http.StatusOK {
		t.Fatalf(
			"http status is %v and body is: %s",
			req.Code, string(req.Body.Bytes()),
		)
	}

	return req.Body.Bytes()
}
