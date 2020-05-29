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

	w.Write([]byte(fmt.Sprintf("Hello, request was made to `%s`\n\n", r.RequestURI)))
	w.Write([]byte("Environment variables:\n"))

	for _, env := range os.Environ() {
		w.Write([]byte(fmt.Sprintf("\t%s\n", env)))
	}
}
