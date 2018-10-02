package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", HelloHandler).Methods("GET")
	r.HandleFunc("/echo", EchoHandler).Methods("GET")
	r.HandleFunc("/incr", IncrHandler).Methods("POST")
	r.HandleFunc("/boom", BoomHandler).Methods("GET")

	srv := &http.Server{
		Handler:      handlers.RecoveryHandler()(r),
		Addr:         "0.0.0.0:8085",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	if text, ok := r.URL.Query()["text"]; ok {
		fmt.Fprintf(w, "%s\n", text[0])
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Missing text param\n")
	}
}

var counter = 0

func IncrHandler(w http.ResponseWriter, r *http.Request) {
	counter += 1
	fmt.Fprintf(w, "%d\n", counter)
}

func BoomHandler(w http.ResponseWriter, r *http.Request) {
	panic("Boom")
}
