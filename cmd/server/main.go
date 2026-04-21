package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/verify", verifyHandler)

	log.Println("Auth Service starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello world to the auth service!"))
	w.WriteHeader(http.StatusOK)
}
