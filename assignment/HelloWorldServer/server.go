package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server listening at localhost:8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
