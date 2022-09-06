package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!, Version = %s, Secret = %s", r.URL.Path[1:], os.Getenv("APP_VERSION"), os.Getenv("APP_SECRET"))
}

func successHealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "Success")
}

func failedHealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	fmt.Fprintf(w, "Failed")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/okhealth", successHealthCheckHandler)
	http.HandleFunc("/nokhealth", failedHealthCheckHandler)
	log.Println("Running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
