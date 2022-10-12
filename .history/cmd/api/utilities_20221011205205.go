package main

import "net/http"

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err
}