package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

type BuildTarget struct {
	Path string `json:"Path"`
	Exec string `json:"Exec"`
	Args string `json:"Args"`
}

func HandlerBuildImpl(w http.ResponseWriter, r *http.Request) {
	var target BuildTarget
	err := json.NewDecoder(r.Body).Decode(&target)
	if err != nil {
		log.Println("cant decode body")
	}
	output := "in `" + target.Path + "` run `" + target.Exec + "` with args: `" + target.Args + "`"

	result := ExecBuildCommand(target.Path, target.Exec, target.Args)
	log.Println(output)
	log.Println(result)
	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte(output + "\n" + "Output :: " + result))
}

func ExecBuildCommand(path, exe, args string) string {
	cmd := exec.Command(fmt.Sprintf(".%v%v", path, exe), args)
	out, err := cmd.Output()
	if err != nil {
		return err.Error()
	}
	return string(out)
}
