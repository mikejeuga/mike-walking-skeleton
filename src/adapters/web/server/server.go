package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
}

func NewServer() *http.Server {
	router := mux.NewRouter()
	s := Server{}

	router.HandleFunc("/", s.Home)

	return &http.Server{
		Addr:    ":8900",
		Handler: router,
	}
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am the skeleton walking!")
}
