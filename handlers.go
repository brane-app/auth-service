package main

import (
	"github.com/gastrodon/groudon/v2"

	"os"
)

var (
	prefix = os.Getenv("PATH_PREFIX")

	routeRoot = "^" + prefix + "/?$"
)

func register_handlers() {
	groudon.AddHandler("POST", routeRoot, postAuth)
}
