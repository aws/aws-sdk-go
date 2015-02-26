package xmlutil

import (
	"encoding/xml"
	"io"
)

type xmlNode struct {
	children   map[string][]*xmlNode
	text       string
	attributes []xml.Attr
}

func xmlToStruct(d *xml.Decoder, s *xml.StartElement) (*xmlNode, error) {
	out := &xmlNode{}
	for {
		tok, err := d.Token()
		if tok == nil || err == io.EOF {
			break
		}
		if err != nil {
			return out, err
		}

		switch typed := tok.(type) {
		case xml.CharData:
			out.text = string(typed.Copy())
		case xml.StartElement:
			if out.children == nil {
				out.children = map[string][]*xmlNode{}
			}

			slice := out.children[typed.Name.Local]
			if slice == nil {
				slice = []*xmlNode{}
			}
			el := typed.Copy()
			node, e := xmlToStruct(d, &el)
			if e != nil {
				return out, e
			}
			slice = append(slice, node)
			out.children[typed.Name.Local] = slice
		case xml.EndElement:
			if s != nil && s.Name.Local == typed.Name.Local { // matching end token
				return out, nil
			}
		}
	}
	return out, nil
}
