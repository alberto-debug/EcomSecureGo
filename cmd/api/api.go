package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/alberto-debug/EcomSecureGo/service/user"
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

	userHandler := user.NEwHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
