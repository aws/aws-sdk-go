package smithytesting

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/internal/smithytesting/xml"
)

// XMLEqual asserts two xml documents by sorting the XML and comparing the strings
// It returns an error in case of mismatch or in case of malformed xml found while sorting.
// In case of mismatched XML, the error string will contain the diff between the two XMLs.
func XMLEqual(expectBytes, actualBytes []byte) error {
	actualString, err := xml.SortXML(bytes.NewBuffer(actualBytes), true)
	if err != nil {
		return err
	}

	expectString, err := xml.SortXML(bytes.NewBuffer(expectBytes), true)
	if err != nil {
		return err
	}

	if !reflect.DeepEqual(expectString, actualString) {
		return fmt.Errorf("XML mismatch\nexpect: %+v\nactual: %+v\n", expectString, actualString)
	}

	return nil
}
