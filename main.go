package main

import (
	"fmt"
	"koriebruh/go_sister/conf"
	"log"
	"net/http"
)

var config *conf.Config

func main() {

	// <-- LOAD .ENV
	config = conf.NewConf()

	// <-- SET HANDLER
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandlerRoot)
	mux.HandleFunc("/home", HandlerHome)

	log.Println("success run in port ", config.Port)
	log.Println("now u connect in server = ", config.Node)
	err := http.ListenAndServe(config.Port, mux)
	if err != nil {
		log.Fatal("server die bg")
	}
}

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome bg ke home")
}

func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "you connect in  : ", config.Node)
}
