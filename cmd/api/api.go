package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NEWAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.Path("/api/v1").Subrouter()

	fmt.Println("Alberto junior")

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
