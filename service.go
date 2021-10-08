package service

import (
	"movieproj/entities"
	Repository "movieproj/repository"

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
func (s Service) AddMovie(m entities.Movies) error {
	m.ID = uuid.New().String()

	err := s.Repo.CreateNewMovie(m)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) GetMovieById(id string) (entities.Movies, error) {
	movie, err := s.Repo.GetMovieById(id)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func (s Service) DeleteMovie(id string) error {
	err := s.DeleteMovie(id)
	if err != nil {
		return err
	}
	return nil
}


