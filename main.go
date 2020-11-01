package main

import (
	"fmt"
	"log"
	"net/http"
)

func Helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Stranger")
	fmt.Println("Endpoint Hit: Helloworld")
}

func handleRequests() {
	http.HandleFunc("/helloworld", Helloworld)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func main() {
	handleRequests()
}
