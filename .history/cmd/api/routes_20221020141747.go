package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func (app *application) routes() http.Hand {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.getOneMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMovies)

	return router
}
