package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {

}