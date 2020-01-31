package main

import (
	"fmt"
	"net/http"
)

func (app *application) ping(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "OK!")
}
