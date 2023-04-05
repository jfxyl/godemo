package main

import (
	"fmt"
	"net/http"
)

func main() {
	var (
		err error
	)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})
	if err = http.ListenAndServe(":1111", nil); err != nil {
		panic(err)
	}
}
