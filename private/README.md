## AWS SDK for Go Private packages ##
`private` is a collection of packages used internally by the SDK, and is subject to have breaking changes at any time. This package is not `internal` so that if need to use this functionality, and understand breaking changes will be made, you are able to.

This package will be refactored in the future so that the API generator and model parsers are exposed cleaning with their own API making it easier for you to auto-generate your own code based on the API models.