package main

import (
	"github.com/omeid/uconfig"
)

func main() {
	uconfig.Classic(&config, nil)
	server := getHTTPServer()
	server.ListenAndServe()
}
