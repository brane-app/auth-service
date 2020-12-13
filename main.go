package main

import (
	"github.com/gastrodon/groudon"
	"git.gastrodon.io/imonke/monkebase"

	"log"
	"net/http"
	"os"
)

func main() {
	monkebase.Connect(os.Getenv("MONKEBASE_CONNECTION"))
	groudon.RegisterHandler("POST", "^/$", postAuth)
	http.Handle("/", http.HandlerFunc(groudon.Route))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
