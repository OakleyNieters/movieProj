package handler

import (
	"github.com/gorilla/mux"
)

type Server struct {
}

func ConfigureRouter(handler MovieHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/movie", handler.PostNewMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", handler.PostNewMovie).Methods("POST")
	r.HandleFunc("/movie", handler.getMovies).Methods("Get")

	return r

}
