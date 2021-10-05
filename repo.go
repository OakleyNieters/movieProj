package Repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"movieproj/entities"
)

type movieStruct struct {
	Movie []entities.Movies
}

type Repo struct {
	Filename string
}

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
	err = json.Unmarshal(bs, ms)
	if err != nil {
		return err
	}
	ms.Movie = append(ms.Movie, mv)

	byteSlice, err := json.Marshal(ms)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, byteSlice, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (r Repo) FindById(id string) (entities.Movies, error) {
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		fmt.Println(err)
	}
	movies := movieStruct{}
	err = json.Unmarshal(file, &movies)
	if err != nil {

		fmt.Println(err)
	}

	moviedb := entities.Movies{}

	for _, movie := range movies.Movie {
		if movie.ID == id {
			moviedb = movie
			return moviedb, nil
		}
	}
	return moviedb, nil

}

func (r Repo) GetMovies() (movieStruct, error) {
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		fmt.Println(err)
	}

	movies := movieStruct{}
	err = json.Unmarshal(file, &movies)
	if err != nil {
		return movies, err
	}
	return movies, nil
}

func (r Repo) DeleteMovie(id string) (entities.Movies, error) {
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		fmt.Println(err)
	}
	movies := movieStruct{}
	err = json.Unmarshal(file, &movies)
	if err != nil {

		fmt.Println(err)
	}

	moviedb := entities.Movies{}

	for index, movie := range movies.Movie {
		if movie.ID == id {
			movies.Movie = append(movies.Movie[:index], movies.Movie[index+1:]...)
			return moviedb, nil
		}
	}
	return moviedb, nil

}

