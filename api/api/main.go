package main

import (
	"net/http"
	"jamagotchi/routes"
)

func main() {
	r := routes.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
