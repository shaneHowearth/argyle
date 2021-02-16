// Package main -
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Routes -
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,    // Log API request calls
		middleware.Recoverer, // Recover from panics without crashing server
	)

	racingRoutes(router)
	return router
}

func main() {
	router := Routes()

	// Walk all the routes and log them
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("Method: %s Route: %s\n", method, route) // Walk and print out all routes
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Fatalf("Logging err: %s\n", err.Error()) // panic if there is an error
	}

	portNum := ""
	ok := false
	if portNum, ok = os.LookupEnv("PORT_NUM"); !ok {
		log.Fatalf("PORT_NUM not set in environment, cannot continue")
	}

	server := &http.Server{Addr: "0.0.0.0:" + portNum, Handler: router}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			// this will always trigger when a shutdown is called
			log.Fatalf("Listen and serve returned error: %v", err)
		}
	}()

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (pkill -2)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown returned error %T", err)
	}
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"message": msg})
}

// respondwithJSON write json response format
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		// log the error
		log.Printf("writing response generated error: %v", err)
	}
}
