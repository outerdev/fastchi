package main

import (
	. "fmt"

	"net/http"

	"github.com/outerdev/fastchi/router"
)

func main() {
	r := router.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request, ps router.Params) {
		Fprintf(w, "Welcome!\n")
	})

	http.ListenAndServe(":8080", r)
}
