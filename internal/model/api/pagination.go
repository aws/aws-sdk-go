package api

import (
	"encoding/json"
	"os"
)

type Paginator struct {
	InputToken  string `json:"input_token"`
	OutputToken string `json:"output_token"`
	LimitKey    string `json:"limit_key"`
	MoreResults string `json:"more_results"`
}

// used for unmarshaling from the paginators JSON file
type paginationDefinitions struct {
	*API
	Pagination map[string]Paginator
}

func (a *API) AttachPaginators(filename string) {
	p := paginationDefinitions{API: a}

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	err = json.NewDecoder(f).Decode(&p)
	if err != nil {
		panic(err)
	}

	p.setup()
}

func (p *paginationDefinitions) setup() {
	for n, e := range p.Pagination {
		paginator := e
		n = p.ExportableName(n)
		if o, ok := p.Operations[n]; ok {
			o.Paginator = &paginator
		} else {
			panic("unknown operation for paginator " + n)
		}
	}
}
