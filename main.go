package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Stranger")
	fmt.Println("Endpoint Hit: Helloworld")
}

func handleRequests() {
	//Creates a new instance of a mux router
	pRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc
	pRouter.HandleFunc("/helloworld", Helloworld)
	log.Fatal(http.ListenAndServe(":8080", pRouter))

}

func main() {
	handleRequests()
}
