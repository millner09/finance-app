package main

import (
	"fmt"
	"net/http"
)

func (app *application) helloFinanceAppHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, finance app")
}
