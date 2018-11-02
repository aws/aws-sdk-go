// +build codegen

package api

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// APIs provides a set of API models loaded by API package name.
type APIs map[string]*API

// LoadAPIs loads the API model files from disk returning the map of API
// package. Returns error if multiple API model resolve to the same package
// name.
func LoadAPIs(modelPaths []string, baseImport string) (APIs, error) {
	apis := APIs{}
	for _, modelPath := range modelPaths {
		a, err := loadAPI(modelPath, baseImport)
		if err != nil {
			return nil, fmt.Errorf("failed to load API, %v, %v", modelPath, err)
		}
		importPath := a.ImportPath()
		if _, ok := apis[importPath]; ok {
			return nil, fmt.Errorf(
				"package names must be unique attempted to load %v twice. Second model file: %v",
				importPath, modelPath)
		}
		apis[importPath] = a
	}

	return apis, nil
}

func loadAPI(modelPath, baseImport string) (*API, error) {
	a := &API{
		BaseImportPath:   baseImport,
		BaseCrosslinkURL: "https://docs.aws.amazon.com",
	}

	err := attachModelFiles(modelPath,
		modelLoader{"api-2.json", a.Attach, true},
		modelLoader{"docs-2.json", a.AttachDocs, true},
		modelLoader{"paginators-1.json", a.AttachPaginators, false},
		modelLoader{"waiters-2.json", a.AttachWaiters, false},
		modelLoader{"examples-1.json", a.AttachExamples, false},
	)
	if err != nil {
		return nil, err
	}

	a.Setup()

	return a, nil
}

type modelLoader struct {
	Filename string
	Loader   func(string)
	Required bool
}

func attachModelFiles(modelPath string, modelFiles ...modelLoader) error {
	for _, m := range modelFiles {
		filepath := filepath.Join(modelPath, m.Filename)
		_, err := os.Stat(filepath)
		if os.IsNotExist(err) && !m.Required {
			continue
		} else if err != nil {
			return fmt.Errorf("failed to load model file %v, %v", m.Filename, err)
		}

		m.Loader(filepath)
	}

	return nil
}

// ExpandModelGlobPath returns a slice of model paths expanded from the glob
// pattern passed in. Returns the directory of the models to be loaded, not the
// path of the model file. Only the most recent service model versions will be
// returned.
//
//   e.g:
//   models/apis/*/*/api-2.json
func ExpandModelGlobPath(globs ...string) ([]string, error) {
	modelPaths := []string{}

	for _, g := range globs {
		filepaths, err := filepath.Glob(g)
		if err != nil {
			return nil, err
		}
		for _, p := range filepaths {
			modelPaths = append(modelPaths, filepath.Dir(p))
		}
	}

	return modelPaths, nil
}

func trimModelServiceVersions(modelPaths []string) []string {
	sort.Strings(modelPaths)

	// Remove old API versions from list
	m := map[string]struct{}{}
	for i := len(modelPaths) - 1; i >= 0; i-- {
		// service name is 2nd-to-last component
		parts := strings.Split(modelPaths[i], string(filepath.Separator))
		svc := parts[len(parts)-2]

		if _, ok := m[svc]; ok {
			// Removed unused service version
			modelPaths = append(modelPaths[:i], modelPaths[i+1:]...)
			continue
		}
		m[svc] = struct{}{}
	}

	return modelPaths
}

// Attach opens a file by name, and unmarshal its JSON data.
// Will proceed to setup the API if not already done so.
func (a *API) Attach(filename string) {
	a.path = filepath.Dir(filename)
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(f).Decode(a); err != nil {
		panic(fmt.Errorf("failed to decode %s, err: %v", filename, err))
	}
}

// AttachString will unmarshal a raw JSON string, and setup the
// API if not already done so.
func (a *API) AttachString(str string) {
	json.Unmarshal([]byte(str), a)

	if !a.initialized {
		a.Setup()
	}
}

// Setup initializes the API.
func (a *API) Setup() {
	a.setServiceAliaseName()
	a.setMetadataEndpointsKey()
	a.writeShapeNames()
	a.resolveReferences()
	a.fixStutterNames()
	a.renameExportable()
	a.applyShapeNameAliases()
	a.createInputOutputShapes()
	a.renameAPIPayloadShapes()
	a.renameCollidingFields()
	a.updateTopLevelShapeReferences()
	a.suppressHTTP2EventStreams()
	a.setupEventStreams()
	a.findEndpointDiscoveryOp()
	a.customizationPasses()

	if !a.NoRemoveUnusedShapes {
		a.removeUnusedShapes()
	}

	if !a.NoValidataShapeMethods {
		a.addShapeValidations()
	}

	a.initialized = true
}
