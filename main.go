package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yashdalfthegray/hue-remote-v2/handlerfuncs"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlerfuncs.Status).Methods("GET")

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	http.Handle("/", loggedRouter)
	fmt.Println("Starting server at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
