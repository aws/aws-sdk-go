# aws-go

aws-go is a set of clients for all Amazon Web Services APIs,
automatically generated from the JSON schemas shipped with
[botocore](http://github.com/boto/botocore).

It supports all known AWS services, and maps exactly to the documented
APIs, with some allowances for Go-specific idioms (e.g. `ID` vs. `Id`).

It is currently **highly untested**, so please be patient and report any
bugs or problems you experience.
