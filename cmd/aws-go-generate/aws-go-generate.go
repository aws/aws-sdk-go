package main

import (
	"os"

	"github.com/stripe/aws-go/aws/model"
)

func main() {
	f, err := os.Open(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	service, err := model.Parse(os.Args[1], f)
	if err != nil {
		panic(err)
	}

	if err := model.Generate(service, os.Stdout); err != nil {
		panic(err)
	}
}
