package main

import (
	"html/template"
	"log"
	"net/http"
)

type Foo struct {
	Name      string
	StaticURL string
	Things    []*Thing
}

type Thing struct {
	Bleep string
	Bloop int
}

func loadTemplates() (*template.Template, error) {
	return template.New("root").ParseGlob("web/*")
}

type Server struct {
	Data      *Foo
	templates *template.Template
}

func NewServer(data *Foo) (*Server, error) {
	templates, err := loadTemplates()
	if err != nil {
		return nil, err
	}

	return &Server{
		Data:      data,
		templates: templates,
	}, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := s.templates.ExecuteTemplate(w, "epicPage", s.Data)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func main() {
	foo := &Foo{
		Name:      "foo bar baz",
		StaticURL: "http://localhost:8000",
		Things: []*Thing{
			&Thing{
				Bleep: "gosho",
				Bloop: 5,
			},
			&Thing{
				Bleep: "tosho",
				Bloop: 42,
			},
		},
	}

	server, err := NewServer(foo)
	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":8080", server))
}
