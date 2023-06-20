package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

//go:embed static
var resources embed.FS

func main() {
	directory, _ := fs.Sub(resources, "static")
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.HandleFunc("/", HttpHandler)
	mux.Handle("/hello/", http.StripPrefix("/hello", fileServer))

	server := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
