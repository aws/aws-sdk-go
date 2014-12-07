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

	if err := model.Load(os.Args[1], f); err != nil {
		panic(err)
	}

	if err := model.Generate(os.Stdout); err != nil {
		panic(err)
	}
}
