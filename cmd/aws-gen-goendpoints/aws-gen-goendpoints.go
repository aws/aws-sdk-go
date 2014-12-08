package main

import (
	"os"

	"github.com/stripe/aws-go/model"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var endpoints model.Endpoints
	if err := endpoints.Parse(f); err != nil {
		panic(err)
	}

	if err := endpoints.Render(os.Stdout); err != nil {
		panic(err)
	}
}
