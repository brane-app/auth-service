package main

import (
	"github.com/gastrodon/groudon/v2"

	"os"
)

func register_handlers() {
	prefix := os.Getenv("PATH_PREFIX")

	groudon.AddHandler("POST", "^"+prefix+"/$", postAuth)
}
