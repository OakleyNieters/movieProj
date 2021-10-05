package service

import (
	"errors"
	"movieproj/Repository"
	"movieproj/entities"

	"github.com/google/uuid"
)

type Service struct {
	Repo Repository.Repo
}

func NewService(r Repository.Repo) Service {
	return Service{
		Repo: r,
	}
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

func (s Service) FindById(id string) (entities.Movies, error) {
	movie, err := s.Repo.FindById(id)
	if err != nil {
		return movie, nil
	}
	return movie, nil
}

func (s Service) DeleteMovie(id string) (entities.Movies, error) {
	movie, err := s.Repo.FindById(id)
	if err != nil {
		return movie, nil
	}
	return movie, nil
}

