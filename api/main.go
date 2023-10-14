package main

import (
	"net/http"
	"api/routes"
)

func main() {
	r := routes.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
