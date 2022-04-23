package handler

import (
	"log"
	"net/http"
)

func CommonHandler() {
	http.HandleFunc("/build", HandlerBuild)
	http.HandleFunc("/test", HandlerTest)
	http.HandleFunc("/add", HandlerAdd)
	http.HandleFunc("/lookup", HandlerLookup)
}

func HandlerBuild(w http.ResponseWriter, r *http.Request) {
	log.Println("handler build")
}

func HandlerTest(w http.ResponseWriter, r *http.Request) {
	log.Println("handler test")
}

func HandlerAdd(w http.ResponseWriter, r *http.Request) {
	log.Println("handler add")
}

func HandlerLookup(w http.ResponseWriter, r *http.Request) {
	log.Println("handler lookup")
}
