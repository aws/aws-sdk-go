// Command aws-gen-gocli parses a JSON description of an AWS API and generates a
// Go file containing a client for the API.
//
//     aws-gen-gocli EC2 apis/ec2/2014-10-01.api.json service/ec2/ec2.go
package main

import (
	"fmt"
	"os"

	"github.com/awslabs/aws-sdk-go/model"
)

func main() {
	in, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer in.Close()

	if err := model.Load(in); err != nil {
		panic(err)
	}

	if err := model.Generate("service"); err != nil {
		fmt.Fprintf(os.Stderr, "error generating %s\n", os.Args[3])
		panic(err)
	}
}
