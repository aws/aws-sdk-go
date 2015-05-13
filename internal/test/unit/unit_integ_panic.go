// Integration version of the unit.go which will panic if imported into and
// a test file while the "integration" tag is set.
//
// +build integration

package unit

// Imported is a marker to ensure that this package's init() function gets
// executed.
//
// To use this package, import it and add:
//
// 	 var _ = integration.Imported
const Imported = true

func init() {
	panic("/internal/test/unit cannot be used with integration testing.")
}