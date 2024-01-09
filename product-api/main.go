package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// reqeusts to the path /goodbye with be handled by this function
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	// any other request will be handled by this function
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Running Hello Handler")

		// read the body
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading body", err)

			http.Error(rw, "Unable to read request body", http.StatusBadRequest)
			return
		}

		// write the response
		fmt.Fprintf(rw, "Hello %s", b)
	})

	fmt.Println("Starting server...")
	// Listen for connections on all ip addresses (0.0.0.0) port 9090
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalf("Unable to start server on port 9090")
	}
}
