package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readNameParam(r *http.Request) (string, error) {
	params := httprouter.ParamsFromContext(r.Context())

	res := params.ByName("name")

	return res, nil
}
