package main

import (
	"CustomDeployer/libs/handler"
	"log"
	"net/http"
	"strconv"
)

const (
	PORT = 1488
)

func main() {
	log.Println("Started program")

	handler.CommonHandler()

	var err = http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
	if err != nil {
		log.Panicln("Server failed starting. Error: %s", err)
	}
}
