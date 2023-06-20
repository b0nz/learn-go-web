package main

import (
	"fmt"
	"net/http"
)

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HttpHandler)

	server := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
