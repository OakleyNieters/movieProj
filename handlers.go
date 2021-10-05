package handler

import (
	"encoding/json"
	"movieproj/entities"
	"movieproj/service"
	"net/http"

	"github.com/gorilla/mux"
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

func (mh MovieHandler) GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	mvID, err := mh.Svc.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	movie, err := json.MarshalIndent(mvID, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(movie)
}

func (mh MovieHandler) DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	mvID, err := mh.Svc.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	movie, err := json.MarshalIndent(mvID, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(movie)
	json.NewEncoder(w).Encode(movie)
}
