package main

import (
	"fmt"
	"net/http"
)

func (app *application) helloFinanceAppHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, finance app")
}

func (app *application) helloParam(w http.ResponseWriter, r *http.Request) {
	res, err := app.readNameParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "name: %s", res)
}
