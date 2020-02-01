package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yashdalfthegray/hue-remote-v2/utils"

	"github.com/yashdalfthegray/hue-remote-v2/handlers"

	gHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if err := utils.CheckEnv(nil); err != nil {
		fmt.Println(err.Error())
		fmt.Println("Exiting with error code 1...")
		os.Exit(1)
	}

	port := utils.GetPort(nil)

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Status).Methods("GET")

	loggedRouter := gHandlers.LoggingHandler(os.Stdout, r)

	http.Handle("/", loggedRouter)
	fmt.Println("Starting server at " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
