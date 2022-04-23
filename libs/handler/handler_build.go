package handler

import (
	"log"
	"net/http"
	"os/exec"
)

func HandlerBuildImpl(w http.ResponseWriter, r *http.Request) {
	execCommand("chmod", "+x", "script")
}

func execCommand(com string, arg ...string) {
	cmd := exec.Command("./bin/script.sh")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(out))
}
