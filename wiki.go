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

func main() {
	http.HandleFunc("/", handler)
	log.Println("Running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
