package handler

import (
	"encoding/json"
	"movieproj/entities"
	"movieproj/service"
	"net/http"
)

type MovieHandler struct {
	Svc service.Service
}

func NewMovieHandler(s service.Service) MovieHandler {
	return MovieHandler{
		Svc: s,
	}
}

func (mh MovieHandler) PostNewMovie(w http.ResponseWriter, r *http.Request) {
	mv := entities.Movies{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = mh.Svc.CreateNewMovie(mv)
	if err != nil {
		switch err.Error() {
		case "invalid movie":
			http.Error(w, err.Error(), http.StatusBadRequest)
		case "movie does not exist":
			http.Error(w, err.Error(), http.StatusNotFound)

		}

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (mh MovieHandler) getMovies(w http.ResponseWriter, r *http.Request) {
	mv := entities.Movies{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = mh.Svc.CreateNewMovie(mv)
	if err != nil {
		switch err.Error() {
		case "invalid movie":
			http.Error(w, err.Error(), http.StatusBadRequest)
		case "movie does not exist":
			http.Error(w, err.Error(), http.StatusNotFound)

		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

	}
}
