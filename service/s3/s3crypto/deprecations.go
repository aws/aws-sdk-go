package s3crypto

import "fmt"

var errDeprecatedCipherBuilder = fmt.Errorf("attempted to use deprecated cipher builder")
var errDeprecatedCipherDataGenerator = fmt.Errorf("attempted to use deprecated cipher data generator")

type deprecatedFeatures interface {
	isUsingDeprecatedFeatures() error
}
