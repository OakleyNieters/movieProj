package repository

import (
	"encoding/json"
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

func getMovies() (Repo, error) {
	file, _ := ioutil.ReadFile("moviesdb.json")
	data := Repo{}
	error := json.Unmarshal([]byte(file), &data)
	if error != nil {
		panic(error)
	}
	return data, error
}

func (r Repo) getMovies(id string) (entities.Movies, error){
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil{
		fmt.Println(err)
	}
	
	movies := DataBase{
		err = json.Unmarshal(file, movies)

		match := entities.Movies{}

		for _, v := range movies.Movies {
			if v.Id == id {
				match = val
				return match, nil
			}
		}
		return entities.Movie{}, errors.New("not found")
	}
}
