package handler

import (
	"github.com/gorilla/mux"
)

type Server struct {
}

func ConfigureRouter(handler MovieHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/movie", handler.PostNewMovie).Methods("POST")

	r.HandleFunc("/movie/{ID}", handler.GetById).Methods("GET")

	r.HandleFunc("/movie/delete/{ID}", handler.DeleteMovieById).Methods("DELETE")

	r.HandleFunc("/movie/{ID}", handler.UpdateMovie).Methods("PUT")

	return r

}
