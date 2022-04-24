package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func HandlerBuildImpl(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
	}

	output := ExecBuildCommand(string(data))
	log.Println(output)

	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte(output))
}

func ExecBuildCommand(arg string) string {
	cmd := exec.Command(fmt.Sprintf("./bin/%v", string(arg)))
	out, err := cmd.Output()
	if err != nil {
		return err.Error()
	}
	return string(out)
}
