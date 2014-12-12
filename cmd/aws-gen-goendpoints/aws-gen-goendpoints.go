package main

import (
	"os"

	"github.com/stripe/aws-go/model"
)

func main() {
	in, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer in.Close()

	var endpoints model.Endpoints
	if err := endpoints.Parse(in); err != nil {
		panic(err)
	}

	out, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer out.Close()

	if err := endpoints.Render(out); err != nil {
		panic(err)
	}
}
