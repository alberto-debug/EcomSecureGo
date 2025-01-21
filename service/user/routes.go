package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{}

func NEwHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handlerLogin).Methods("POST")
	router.HandleFunc("/register", h.handlerLogin).Methods("POST")
}

func (h *Handler) handlerLogin(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) handlerRegister(w http.ResponseWriter, r *http.Request) {
}
