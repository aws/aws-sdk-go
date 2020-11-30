// +build gofuzz

package xmlutil

import (
	"bytes"
	"encoding/xml"
)

type mockOutput struct {
	_       struct{}          `type:"structure"`
	String  *string           `type:"string"`
	Integer *int64            `type:"integer"`
	Nested  *mockNestedStruct `type:"structure"`
	List    []*mockListElem   `locationName:"List" locationNameList:"Elem" type:"list"`
	Closed  *mockClosedTags   `type:"structure"`
}
type mockNestedStruct struct {
	_            struct{} `type:"structure"`
	NestedString *string  `type:"string"`
	NestedInt    *int64   `type:"integer"`
}
type mockClosedTags struct {
	_    struct{} `type:"structure" xmlPrefix:"xsi" xmlURI:"http://www.w3.org/2001/XMLSchema-instance"`
	Attr *string  `locationName:"xsi:attrval" type:"string" xmlAttribute:"true"`
}
type mockListElem struct {
	_          struct{}            `type:"structure" xmlPrefix:"xsi" xmlURI:"http://www.w3.org/2001/XMLSchema-instance"`
	String     *string             `type:"string"`
	NestedElem *mockNestedListElem `type:"structure"`
}
type mockNestedListElem struct {
	_ struct{} `type:"structure" xmlPrefix:"xsi" xmlURI:"http://www.w3.org/2001/XMLSchema-instance"`

	String *string `type:"string"`
	Type   *string `locationName:"xsi:type" type:"string" xmlAttribute:"true"`
}

func Fuzz(data []byte) int {
	actual := mockOutput{}
	decoder := xml.NewDecoder(bytes.NewReader(data))
	err := UnmarshalXML(&actual, decoder, "")
	if err != nil {
		return 0
	}
	return 1
}
