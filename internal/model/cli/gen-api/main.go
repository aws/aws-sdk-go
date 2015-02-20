// Command aws-gen-gocli parses a JSON description of an AWS API and generates a
// Go file containing a client for the API.
//
//     aws-gen-gocli apis/s3/2006-03-03.normal.json
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"

	"github.com/awslabs/aws-sdk-go/internal/model/api"
)

type generateInfo struct {
	*api.API
	ForceService bool
	PackageDir   string
}

func newGenerateInfo(modelFile, svcPath string, forceService bool) *generateInfo {
	g := &generateInfo{API: &api.API{}, ForceService: forceService}
	g.API.Attach(modelFile)

	// ensure the directory exists
	pkgDir := filepath.Join(svcPath, g.API.PackageName())
	os.MkdirAll(pkgDir, 0775)

	g.PackageDir = pkgDir

	return g
}

func main() {
	var svcPath string
	var forceService bool
	flag.StringVar(&svcPath, "path", "service", "generate in a specific directory (default: 'service')")
	flag.BoolVar(&forceService, "force", false, "force re-generation of PACKAGE/service.go")
	flag.Parse()

	files := []string{}
	for i := 0; i < flag.NArg(); i++ {
		file := flag.Arg(i)
		if strings.Contains(file, "*") {
			paths, _ := filepath.Glob(file)
			files = append(files, paths...)
		} else {
			files = append(files, file)
		}
	}

	for _, file := range files {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmtStr := "Error generating %s\n%s\n%s\n"
					fmt.Fprintf(os.Stderr, fmtStr, file, r, debug.Stack())
				}
			}()

			g := newGenerateInfo(file, svcPath, forceService)

			// write api.go and service.go files
			g.writeAPIFile()
			g.writeServiceFile()
		}()
	}
}

func (g *generateInfo) writeServiceFile() {
	file := filepath.Join(g.PackageDir, "service.go")
	if _, err := os.Stat(file); g.ForceService || (err != nil && os.IsNotExist(err)) {
		ioutil.WriteFile(file, []byte("package "+g.API.PackageName()+"\n\n"+g.API.ServiceGoCode()), 0664)
	}
}

func (g *generateInfo) writeAPIFile() {
	file := filepath.Join(g.PackageDir, "api.go")
	ioutil.WriteFile(file, []byte("package "+g.API.PackageName()+"\n\n"+g.API.APIGoCode()), 0664)
}
