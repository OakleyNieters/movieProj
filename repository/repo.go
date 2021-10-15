package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"movieproj/entities"
)

type movieStruct struct {
	Movies []entities.Movies
}

type Repo struct {
	Filename string
}

//type Repository interface {
//	GetMovieById(id string) entities.Movies
//}

func NewRepository(filename string) Repo {
	return Repo{
		Filename: filename,
	}
}

func (r Repo) CreateNewMovie(mv entities.Movies) error {
	ms := movieStruct{}

	bs, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bs, &ms)
	if err != nil {
		return err
	}
	for _, val := range ms.Movies {
		if val.Title == mv.Title {
			return errors.New("movie already exists")
		}
	}

	ms.Movies = append(ms.Movies, mv)

	byteSlice, err := json.MarshalIndent(ms, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, byteSlice, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (r Repo) GetMovieById(id string) (entities.Movies, error) {
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		fmt.Println(err)
	}
	movies := movieStruct{}
	err = json.Unmarshal(file, &movies)

	compare := entities.Movies{}

	for _, val := range movies.Movies {
		if val.ID == id {
			compare = val
			return compare, nil

		}

	}
	return compare, nil
}

func (r *Repo) DeleteMovieById(id string) error {
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	movies := movieStruct{}
	newDb := movieStruct{}
	err = json.Unmarshal(file, &movies)

	dbLen := len(movies.Movies)
	for _, val := range movies.Movies {
		if val.ID == id {
			continue
		} else {
			newDb.Movies = append(newDb.Movies, val)
		}
	}

	if len(newDb.Movies) == dbLen {
		return errors.New("failed to delete movie - does not exist")
	}

	movieBytes, err := json.MarshalIndent(newDb, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, movieBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) UpdateById(id string, ms entities.Movies) error {
	movie := movieStruct{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &movie)
	if err != nil {
		return err
	}
	for _, val := range movie.Movies {
		if val.ID == id {
			movie.Movies = append(movie.Movies, ms)

		}
	}
	movieBytes, err := json.MarshalIndent(movie.Movies, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(r.Filename, movieBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
