package service

import (
	"errors"
	"movieproj/entities"
	"movieproj/repository"

	//"movieproj/repository"

	"github.com/google/uuid"
)

type service struct {
	Repo repository.Repo
}

type Repo interface {
	DeleteMovie(id string) error
	UpdateById(id string, ms entities.Movies) error
	AddMovie(entities.Movies) error
	GetMovieById(id string) (entities.Movies, error)
}

func NewService(r repository.Repo) service {
	return service{
		Repo: r,
	}
}
func (s service) AddMovie(m entities.Movies) error {
	m.ID = uuid.New().String()

	err := s.AddMovie(m)
	if err != nil {
		return err
	}
	return nil
}

func (s service) GetMovieById(id string) (*entities.Movies, error) {
	movie, err := s.GetMovieById(id)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func (s service) DeleteMovie(id string) error {
	err := s.DeleteMovie(id)
	if err != nil {
		return err
	}
	return nil
}

func (s service) UpdateById(id string, ms entities.Movies) error {
	if id != ms.ID {
		return errors.New("invalid ID input")
	}
	err := s.UpdateById(id, ms)
	if err != nil {
		return err
	}
	return nil
}
