package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/helloFinanceApp", app.helloFinanceAppHandler)
	router.HandlerFunc(http.MethodGet, "/v1/hello/:name", app.helloParam)

	return router
}
