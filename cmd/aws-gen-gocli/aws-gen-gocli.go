// Command aws-gen-gocli parses a JSON description of an AWS API and generates a
// Go file containing a client for the API.
//
//     aws-gen-gocli EC2 apis/ec2/2014-10-01.api.json service/ec2/ec2.go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/awslabs/aws-sdk-go/model"
)

func main() {
	var svcPath string
	var forceService bool
	flag.StringVar(&svcPath, "path", "service", "generate in a specific directory (default: 'service')")
	flag.BoolVar(&forceService, "force", false, "force re-generation of PACKAGE/service.go")
	flag.Parse()

	api := os.Args[len(os.Args)-flag.NArg()]

	in, err := os.Open(api)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	if err := model.Load(in); err != nil {
		panic(err)
	}

	if err := model.Generate(svcPath, forceService); err != nil {
		fmt.Fprintf(os.Stderr, "error generating %s\n", os.Args[1])
		panic(err)
	}
}
