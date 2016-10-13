package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"os"
	"regexp"
	"strings"

	xhtml "golang.org/x/net/html"
)

type apiDocumentation struct {
	*API
	Operations map[string]string
	Service    string
	Shapes     map[string]shapeDocumentation
}

type shapeDocumentation struct {
	Base string
	Refs map[string]string
}

// AttachDocs attaches documentation from a JSON filename.
func (a *API) AttachDocs(filename string) {
	d := apiDocumentation{API: a}

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	err = json.NewDecoder(f).Decode(&d)
	if err != nil {
		panic(err)
	}

	d.setup()

}

func (d *apiDocumentation) setup() {
	d.API.Documentation = docstring(d.Service)
	if d.Service == "" {
		d.API.Documentation =
			fmt.Sprintf("// %s is a client for %s.\n", d.API.StructName(), d.API.NiceName())
	}

	for op, doc := range d.Operations {
		d.API.Operations[op].Documentation = strings.TrimSpace(docstring(doc))
	}

	for shape, info := range d.Shapes {
		if sh := d.API.Shapes[shape]; sh != nil {
			sh.Documentation = docstring(info.Base)
		}

		for ref, doc := range info.Refs {
			if doc == "" {
				continue
			}

			parts := strings.Split(ref, "$")
			if sh := d.API.Shapes[parts[0]]; sh != nil {
				if m := sh.MemberRefs[parts[1]]; m != nil {
					m.Documentation = docstring(doc)
				}
			}
		}
	}
}

// TODO
// Change to use new docstring generator
var reNewline = regexp.MustCompile(`\r?\n`)
var reMultiSpace = regexp.MustCompile(`\s+`)
var reComments = regexp.MustCompile(`<!--.*?-->`)
var reFullname = regexp.MustCompile(`\s*<fullname?>.+?<\/fullname?>\s*`)
var reExamples = regexp.MustCompile(`<examples?>.+?<\/examples?>`)
var rePara = regexp.MustCompile(`<(?:p|h\d)>(.+?)</(?:p|h\d)>`)
var reLink = regexp.MustCompile(`<a href="(.+?)">(.+?)</a>`)
var reTag = regexp.MustCompile(`<.+?>`)
var reEndNL = regexp.MustCompile(`\n+$`)

// docstring rewrites a string to insert godocs formatting.
func docstring(doc string) string {
	doc = reNewline.ReplaceAllString(doc, "")
	doc = reMultiSpace.ReplaceAllString(doc, " ")
	doc = reComments.ReplaceAllString(doc, "")
	doc = reFullname.ReplaceAllString(doc, "")
	doc = reExamples.ReplaceAllString(doc, "")

	_, err := xhtml.Parse(strings.NewReader(doc))
	// If there is no error, means it is HTML. However, if it isn't HTML, we use the old
	// way of generating documentation.
	if err == nil {
		// this will only occur if an error during the tokenization process occurs.
		doc = generateDocFromHTML(doc)
	} else {
		doc = rePara.ReplaceAllString(doc, "$1\n\n")
		doc = reLink.ReplaceAllString(doc, "$2 ($1)")
	}

	doc = reEndNL.ReplaceAllString(doc, "")
	if doc == "" {
		return "\n"
	}

	doc = html.UnescapeString(doc)
	return commentify(doc)
}

// generateDocFromHTML will generate the proper doc string for html encoded doc entries.
func generateDocFromHTML(htmlSrc string) string {
	tokenizer := xhtml.NewTokenizer(strings.NewReader(htmlSrc))
	// tagInfo contains necessary token info of start tag
	type tagInfo struct {
		tag string
		key string
		val string
		txt string
	}
	doc := ""
	level := 0
	col := 0
	indent := false
	stack := []tagInfo{}

	for tt := tokenizer.Next(); tt != xhtml.ErrorToken; tt = tokenizer.Next() {
		switch tt {
		case xhtml.ErrorToken:
			break
		case xhtml.TextToken:
			txt := string(tokenizer.Text())
			if len(stack) > 0 {
				// set the current entries txt value. We use this
				// for checking if we need to indent. The indent rules
				// is if there is nested html tags that aren't empty or space,
				// we will indent the code or li blocks.
				size := len(stack)
				stack[size-1].txt = txt
				info := stack[size-1]

				// if the tag is <code> or <li> we want to indent this block
				if info.tag == "pre" || info.tag == "li" {
					indent = true
				} else if info.tag == "code" {
					if level == 1 {
						indent = true
					} else {
						for i := level; size-i >= 0 && i > 1; i-- {
							if stack[size-i].tag == "li" || stack[size-i].tag == "pre" {
								indent = true
								break
							}
						}

					}
				}

				// we only want to append to the doc if the current txt isn't
				// empty or contains only a space.
				if len(txt) > 0 && txt != " " {
					if info.tag == "a" {
						if info.val != "" {
							txt += fmt.Sprintf(" (%s)", info.val)
						}
					}

					// check spacing
					// we do not care about the current scope we are on, hence the
					// i > 1. The reason to check for the check against one is due to the
					// fact that empty stack is at level 0, which we don't care about.
					for i := level; size-i >= 0 && i > 1; i-- {
						if stack[size-i].tag == "p" {
							if len(stack[size-i].txt) > 0 && stack[size-i].txt != " " {
								indent = false
							}
						}
					}

					txt, col = wrap(txt, col, 72, indent)
					doc += txt
				}
			} else {
				indent = false
				txt, col = wrap(txt, col, 72, indent)
				doc += txt
			}
		case xhtml.StartTagToken:
			tn, _ := tokenizer.TagName()
			key, val, _ := tokenizer.TagAttr()
			info := tagInfo{
				tag: string(tn),
				key: string(key),
				val: string(val),
			}
			stack = append(stack, info)
			level++
		case xhtml.SelfClosingTagToken, xhtml.EndTagToken:
			// the stack could be empty here
			indent = false
			if len(stack) > 0 {
				info := stack[len(stack)-1]
				if info.tag == "p" || (len(info.tag) == 2 && info.tag[0] == 'h') {
					doc += "\n\n"
					col = 0
				} else if (level == 1 && info.tag == "code") || info.tag == "pre" || info.tag == "li" {
					doc += "\n"
					col = 0
				}
				stack = stack[:len(stack)-1]
				level--
			}
		}
	}
	return doc
}

// commentify converts a string to a Go comment
func commentify(doc string) string {
	lines := strings.Split(doc, "\n")
	out := []string{}
	for i, line := range lines {
		if i > 0 && line == "" && lines[i-1] == "" {
			continue
		}
		out = append(out, "// "+line)
	}

	return strings.Join(out, "\n") + "\n"
}

// wrap returns a rewritten version of text to have line breaks
// at approximately length characters. Line breaks will only be
// inserted into whitespace.
func wrap(text string, col, length int, indent bool) (string, int) {
	var buf bytes.Buffer
	var last rune
	var lastNL bool

	if col == 0 && indent {
		buf.WriteString("   ")
		col += 3
	}
	for _, c := range text {
		switch c {
		case '\r': // ignore this
			continue // and also don't track `last`
		case '\n': // ignore this too, but reset col
			if col >= length || last == '\n' {
				buf.WriteString("\n")
			}
			buf.WriteString("\n")
			col = 0
			if indent {
				buf.WriteString("   ")
				col += 3
			}
		case ' ', '\t': // opportunity to split
			if col >= length {
				buf.WriteByte('\n')
				col = 0
				if indent {
					buf.WriteString("   ")
					col += 3
				}
			} else {
				// We only want to write a leading space if the col is greater than zero.
				// This will provide the proper spacing for documentation.
				if !lastNL && col > 0 {
					buf.WriteRune(c)
					col++ // count column
				}
			}
		default:
			buf.WriteRune(c)
			col++
		}
		lastNL = c == '\n'
		last = c
	}
	return buf.String(), col
}
