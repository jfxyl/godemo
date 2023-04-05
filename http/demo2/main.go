package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	var (
		router *mux.Router
	)
	router = mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.ListenAndServe(":3333", router)
}
