package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/verify", verifyHandler)

	log.Println("Auth Service starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Time from verifyHandler: ", time.Now().Format(time.RFC3339))

	// temp to test middleware in k3s
	w.Header().Set("X-User-ID", "777")
	w.Header().Set("X-Auth-Status", "verified")

	w.Write([]byte("Hello world to the auth service!"))
	w.WriteHeader(http.StatusOK)
}
