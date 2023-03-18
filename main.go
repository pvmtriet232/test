package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})

	fmt.Println("Server listening on port 4000...")
	http.ListenAndServe(":4000", nil)
}
