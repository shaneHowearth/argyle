package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/shanehowearth/argyle/fibonacci"
)

func racingRoutes(router *chi.Mux) {
	// Articles related routes
	router.Route("/", func(r chi.Router) {
		r.Get("/fib", GetFibonacci)
	})
}

// GetFibonacci -
func GetFibonacci(w http.ResponseWriter, req *http.Request) {
	nth := req.URL.Query().Get("n")
	if nth == "" {
		log.Printf("Invalid nth string: %s", req.URL.Query().Get("n"))
		respondWithError(w, http.StatusInternalServerError, "Did you supply an n value?")
		return
	}
	var n int
	var err error
	if n, err = strconv.Atoi(nth); err != nil {
		log.Printf("Converting %s to an int resulted in error %v", nth, err)
		// We don't want the user to know about the inner workings of the application
		respondWithError(w, http.StatusInternalServerError, "An invalid n value was supplied, please try again")
		return
	}
	val := fibonacci.Compute(n)
	// Left for debugging purposes
	// log.Print("Calculated fibonacci value: ", val.String())
	_, err = w.Write([]byte(fmt.Sprint(val.String())))
}
