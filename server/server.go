package server

import (
	"fmt"
    "net/http"
)

type Server struct{

}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/get_movies", s.getMovies)

	mux.ServeHTTP(w, r)
}

func (s *Server) getMovies(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "yo yo yo")
}
