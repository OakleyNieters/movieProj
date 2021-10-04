package service

import (
	"errors"
	repository "movieproj/Repository"
	"movieproj/entities"

	"github.com/google/uuid"
)

type Service struct {
	Repo repository.Repo
}

func (s Service) CreateNewMovie(mv entities.Movies) error {

	mv.ID = uuid.New().String()

	if mv.Rating <= 10 && mv.Rating >= 0 {
		err := s.Repo.CreateNewMovie(mv)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("invalid Rating")
}

func (s Service) getMovies(Id string) (entities.Movies, error) {
	movie, err := s.getMovies(Id)
	if err != nil {
		return movie, err
	}
	return movie, nil
}
