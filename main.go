package main

import (
	"fmt"
	"log"
	"net/http"
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
	message := fmt.Sprintf("Hello, request was made to `%s`", r.RequestURI)
	w.Write([]byte(message))
}
