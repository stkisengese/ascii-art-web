package main

import (
	"ascii/handlers"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	fmt.Printf("Starting server on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {

	case "/":
		handlers.IndexHandler(w, r)

	case "/ascii-art":
		handlers.AsciiArtHandler(w, r)

	default:
		http.NotFound(w, r)
	}

}
