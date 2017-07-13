package github

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Commit struct {
	Id  string
	Url string
}

type Webhook struct {
	Ref        string `json:"ref"`
	Id         string `json:"id"`
	Commits    []Commit
	Repository struct {
		Name           string `json:"name"`
		Url            string `json:"url"`
		Default_branch string `json:"default_branch"`
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	secret := []byte(os.Getenv("GITHUB_TOKEN"))
	hook, err := Parse(secret, r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var b bytes.Buffer
	err = json.Indent(&b, hook.Payload, "", "\t")
	if err != nil {
		log.Println("JSON parse error: ", err)
		return
	}

	log.Printf("---\n %s \n", b.Bytes())
}
