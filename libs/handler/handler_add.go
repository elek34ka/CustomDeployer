package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

type Repository struct {
	User     string `json:"User"`
	RepoName string `json:"RepoName"`
}

func HandlerAddImpl(w http.ResponseWriter, r *http.Request) {
	var repo Repository
	err := json.NewDecoder(r.Body).Decode(&repo)
	if err != nil {
		log.Println(err.Error())
		w.Header().Add("Content-type", "text/html")
		w.Write([]byte(err.Error()))
		return
	}

	output := fmt.Sprintf("github.com/%v/%v", repo.User, repo.RepoName)
	result := ExecAddCommand("git_clone.sh", repo.User, repo.RepoName)
	log.Println(output + " " + result)
	w.Header().Add("Content-type", "text/html")
	w.Write([]byte(output + " " + result))
}

func ExecAddCommand(file, user, repoName string) string {
	path := fmt.Sprintf("git/users/%v", user)
	out, err := exec.Command(fmt.Sprintf("./bin/%v", file), user, repoName, path).Output()
	if err != nil {
		return err.Error()
	}
	return string(out) + "All is ok"
}
