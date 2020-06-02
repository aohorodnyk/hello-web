package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerMain)

	host := ":8080"
	log.Printf("Running web application on host:port `%s`\n", host)
	err := http.ListenAndServe(host, mux)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Web application was successfuly finished")
}

func handlerMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintf(w, "Hello, request was made to `%s`\n\n", r.RequestURI)

	getParams := r.URL.Query()
	if len(getParams) > 0 {
		w.Write([]byte("GET params:\n"))
		for getName := range getParams {
			getValue := r.URL.Query().Get(getName)
			fmt.Fprintf(w, "\t%s: %s\n", getName, getValue)
		}

		fmt.Fprintf(w, "\n\n")
	}

	w.Write([]byte("Headers:\n"))
	for headerName := range r.Header {
		headerValue := r.Header.Get(headerName)
		fmt.Fprintf(w, "\t%s: %s\n", headerName, headerValue)
	}

	w.Write([]byte(fmt.Sprint("\n\n")))

	w.Write([]byte("Environment variables:\n"))
	for _, env := range os.Environ() {
		fmt.Fprintf(w, "\t%s\n", env)
	}
}
