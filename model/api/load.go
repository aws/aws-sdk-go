package api

import (
	"encoding/json"
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
	a.writeShapeNames()
	a.resolveReferences()
	a.renameExportable()
	a.renameToplevelShapes()
	a.initialized = true
}
