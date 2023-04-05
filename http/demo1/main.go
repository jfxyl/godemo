package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	var (
		err        error
		httpServer http.Server
		mux        *http.ServeMux
	)
	mux = http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})
	httpServer = http.Server{
		Addr:         ":2222",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	if err = httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}
