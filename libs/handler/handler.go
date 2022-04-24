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
	http.HandleFunc("/login", HandlerLogin)
}

func HandlerBuild(w http.ResponseWriter, r *http.Request) {
	log.Println("handler build")
	HandlerBuildImpl(w, r)
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

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("handler login")
	HandlerLoginImpl(w, r)
}
