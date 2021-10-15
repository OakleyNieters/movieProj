package handler

import (
	"encoding/json"
	"fmt"
	"movieproj/entities"

	//"movieproj/service"

	//"movieproj/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Service interface {
	DeleteMovie(id string) error
	UpdateById(id string, ms entities.Movies) error
	AddMovie(entities.Movies) error
	GetMovieById(id string) (*entities.Movies, error)
	//GetById(w http.ResponseWriter, r *http.Request)
}

type MovieHandler struct {
	Svc Service
}

func NewMovieHandler(s Service) MovieHandler {
	return MovieHandler{
		Svc: s,
	}
}

func (mh MovieHandler) PostNewMovie(w http.ResponseWriter, r *http.Request) {
	mv := entities.Movies{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		fmt.Println(err)
	}

	err = mh.Svc.AddMovie(mv)
	if err != nil {
		switch err.Error() {
		case "movie already exists":
			http.Error(w, err.Error(), http.StatusBadRequest)
		case "invalid rating":
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (mh MovieHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	mvID, err := mh.Svc.GetMovieById(id)
	if err != nil {
		switch err.Error() {
		case "movie not found":
			http.Error(w, err.Error(), http.StatusNotFound)
		}
	}

	movie, err := json.MarshalIndent(mvID, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(movie)
}

func (mov MovieHandler) DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	err := mov.Svc.DeleteMovie(id)
	if err != nil {
		switch err.Error() {
		case "failed to delete movie - does not exist":
			http.Error(w, err.Error(), http.StatusNotFound)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (mh MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	mv := entities.Movies{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = mh.Svc.UpdateById(id, mv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
