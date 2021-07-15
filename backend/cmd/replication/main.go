package main

import (
	"context"

	"github.com/omeid/uconfig"
)

func main() {
	uconfig.Classic(&config, nil)
	worker := getWorker()
	err := worker.Run(context.Background())
	if err != nil {
		panic(err)
	}
}
