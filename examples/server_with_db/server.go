package bleh

import (
	"fmt"
	"net/http"
)

type BlehServer struct {
	bleh *Bleh
}

func NewBlehServer(bleh *Bleh) *BlehServer {
	return &BlehServer{
		bleh: bleh,
	}
}

func (b *BlehServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/foo", b.foo)
	mux.HandleFunc("/get_age", b.getAge)

	mux.ServeHTTP(w, r)
}

func (b *BlehServer) foo(w http.ResponseWriter, r *http.Request) {
	message, err := b.bleh.Foo()

	if err != nil {
		http.Error(w, fmt.Sprintf("Foo failed: %s", err), 500)
		return
	}

	fmt.Fprintf(w, "%s", message)
}

func (b *BlehServer) getAge(w http.ResponseWriter, r *http.Request) {
	person := r.URL.Query().Get("person")

	age, err := b.bleh.GetAge(person)

	if err != nil {
		http.Error(w, fmt.Sprintf("GetAge failed: %s", err), 500)
		return
	}

	fmt.Fprintf(w, "%s's age is %d", person, age)
}
