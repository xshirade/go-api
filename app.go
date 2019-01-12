package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Gorilla!\n"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", router))
}
