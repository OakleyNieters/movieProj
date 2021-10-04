package main

import (
	"log"
	Repository "movieproj/repository"
	"movieproj/service"
	"path/filepath"

	handler "movieproj/handlers"
	"net/http"
)

func main() {

	fileName := "/Users/oakleynieters/Downloads/MovieProj/moviedb.json"

	ext := filepath.Ext(fileName)
	if ext != ".json" {
		log.Fatalln("File invalid")

	}
	r := Repository.NewRepository(fileName)

	svc := service.Service(r)

	hdlr := handler.NewMovieHandler(svc)

	router := handler.ConfigureRouter(hdlr)

	svr := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	log.Fatalln(svr.ListenAndServe())

}
