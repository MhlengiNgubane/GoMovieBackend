package main

import "net/http"

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {

}