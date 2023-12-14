package router

import (
	"github.com/gorilla/mux"
	"lms_try/handler"
	"net/http"
)

func GenerateUserRouter(r *mux.Router, handler *handler.UserHandler) *mux.Router {
	r.HandleFunc("/users", handler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/user", handler.GetById).Methods(http.MethodPost)
	return r
}
