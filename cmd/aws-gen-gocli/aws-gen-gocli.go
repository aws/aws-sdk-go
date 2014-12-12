package main

import (
	"os"

	"github.com/stripe/aws-go/model"
)

func main() {
	in, err := os.Open(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out, err := os.Create(os.Args[3])
	if err != nil {
		panic(err)
	}
	defer out.Close()

	if err := model.Load(os.Args[1], in); err != nil {
		panic(err)
	}

	if err := model.Generate(out); err != nil {
		panic(err)
	}
}
