package api

import (
	"encoding/json"
	"fmt"
	"os"
)

func Load(api, docs, paginators, waiters string) *API {
	a := API{}
	a.Attach(api)
	a.Attach(docs)
	a.Attach(paginators)
	a.Attach(waiters)
	return &a
}

func (a *API) Attach(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	json.NewDecoder(f).Decode(a)

	if !a.initialized {
		a.Setup()
	}
}

func (a *API) AttachString(str string) {
	json.Unmarshal([]byte(str), a)

	if !a.initialized {
		a.Setup()
	}
}

func (a *API) Setup() {
	a.unrecognizedNames = map[string]string{}
	a.writeShapeNames()
	a.resolveReferences()
	a.renameExportable()
	a.renameToplevelShapes()
	a.updateTopLevelShapeReferences()
	a.createInputOutputShapes()

	if !a.NoRemoveUnusedShapes {
		a.removeUnusedShapes()
	}

	if len(a.unrecognizedNames) > 0 {
		msg := []string{
			"Unrecognized inflections for the following export names:",
			"(Add these to inflections.csv with any inflections added after the ':')",
		}
		fmt.Fprintf(os.Stderr, "%s\n%s\n\n", msg[0], msg[1])
		for n, m := range a.unrecognizedNames {
			if n == m {
				m = ""
			}
			fmt.Fprintf(os.Stderr, "%s:%s\n", n, m)
		}
		os.Stderr.WriteString("\n\n")
		panic("Found unrecognized exported names in API " + a.PackageName())
	}

	a.initialized = true
}
