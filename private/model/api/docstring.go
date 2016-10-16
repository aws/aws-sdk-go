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

var reNewline = regexp.MustCompile(`\r?\n`)
var reMultiSpace = regexp.MustCompile(`\s+`)
var reComments = regexp.MustCompile(`<!--.*?-->`)
var reFullname = regexp.MustCompile(`\s*<fullname?>.+?<\/fullname?>\s*`)
var reExamples = regexp.MustCompile(`<examples?>.+?<\/examples?>`)
var reEndNL = regexp.MustCompile(`\n+$`)

// docstring rewrites a string to insert godocs formatting.
func docstring(doc string) string {
	doc = reNewline.ReplaceAllString(doc, "")
	doc = reMultiSpace.ReplaceAllString(doc, " ")
	doc = reComments.ReplaceAllString(doc, "")
	doc = reFullname.ReplaceAllString(doc, "")
	doc = reExamples.ReplaceAllString(doc, "")

	doc = generateDoc(doc)
	doc = reEndNL.ReplaceAllString(doc, "")
	if doc == "" {
		return "\n"
	}

	doc = html.UnescapeString(doc)
	return commentify(doc)
}

// style is what we want to prefix a string with.
// For instance, <li>Foo</li><li>Bar</li>, will generate
//    * Foo
//    * Bar
var style = map[string]string{
	"li": "* ",
}

const (
	indent = "   "
)

// generateDoc will generate the proper doc string for html encoded or plain text doc entries.
/*func generateDoc(htmlSrc string) string {
	tokenizer := xhtml.NewTokenizer(strings.NewReader(htmlSrc))
	// tagInfo contains necessary token info of start tag
	type tagInfo struct {
		tag        string
		key        string
		val        string
		txt        string
		closingTag bool
	}

	doc := ""
	level := 0
	col := 0
	isIndented := true
	alwaysIndent := false
	isInlined := true
	stack := []tagInfo{}
	stringBlock := []tagInfo{}

	for tt := tokenizer.Next(); tt != xhtml.ErrorToken; tt = tokenizer.Next() {
		switch tt {
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
				shouldIndent := false
				shouldInline := false

				// if the tag is <code> or <li> we want to indent this block
				if info.tag == "pre" || info.tag == "li" || info.tag == "ul" {
					alwaysIndent = true
				} else if info.tag == "code" {
					if level == 1 {
						shouldIndent = true
					} else {
						for i := len(stack) - 1; i >= 0 && i > 0; i-- {
							if stack[i].tag == "li" || stack[i].tag == "pre" {
								// We only care about li and pre parent tags
								shouldIndent = true
								break
							}
						}
					}
				} else {
					shouldIndent = false
					shouldInline = true
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
					for i := len(stack) - 1; i >= 0 && i > 1; i-- {
						if stack[i].tag == "p" {
							if len(stack[i].txt) > 0 && stack[i].txt != " " {
								shouldIndent = false
							}
						}
					}

					info.txt = txt
					isIndented = isIndented && shouldIndent
					isInlined = isInlined && shouldInline
				} else {
					info.txt = ""
				}
				stringBlock = append(stringBlock, info)
			} else {
				isIndented = true
				txt, col = wrap(txt, col, 72, false)
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
			if len(stringBlock) > 0 {
				tn, _ := tokenizer.TagName()
				key, val, _ := tokenizer.TagAttr()
				info := tagInfo{
					tag:        string(tn),
					key:        string(key),
					val:        string(val),
					closingTag: true,
				}
				stringBlock = append(stringBlock, info)
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
				level--
				if level == 0 {
					txt := ""
					for i := 0; i < len(stringBlock); i++ {
						info = stringBlock[i]
						if col == 0 && (isIndented || alwaysIndent) {
							if len(stringBlock) > 0 {
								fmt.Println("HERE", info.tag)
								v, ok := style[info.tag]
								if ok && !info.closingTag {
									indentStr := indent
									indentStr += v
									info.txt = indentStr + info.txt
									col = len(indentStr)
									fmt.Println(indentStr)
								}
							}
						}
						if info.txt != "" {
							info.txt, col = wrap(info.txt, col, 72, alwaysIndent || isIndented)
							if info.closingTag {
								if info.tag == "p" || (len(info.tag) == 2 && info.tag[0] == 'h') {
									info.txt += "\n\n"
									col = 0
								} else if (i == 0 && info.tag == "code") || info.tag == "pre" || info.tag == "li" {
									info.txt += "\n"
									col = 0
								}
							}
						}
						txt += info.txt
					}
					doc += txt
					stringBlock = []tagInfo{}
					isIndented = true
					alwaysIndent = false
				}
			}
		}
	}
	return doc
}*/

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
func wrap(text string, col, length int, isIndented bool) (string, int) {
	var buf bytes.Buffer
	var last rune
	var lastNL bool

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
			if isIndented {
				buf.WriteString(indent)
				col += 3
			}
		case ' ', '\t': // opportunity to split
			if col >= length {
				buf.WriteByte('\n')
				col = 0
				if isIndented {
					buf.WriteString(indent)
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

type tagInfo struct {
	tag        string
	key        string
	val        string
	txt        string
	raw        string
	closingTag bool
}

// generateDoc will generate the proper doc string for html encoded or plain text doc entries.
func generateDoc(htmlSrc string) string {
	tokenizer := xhtml.NewTokenizer(strings.NewReader(htmlSrc))
	stack := buildStack(tokenizer)
	return walk(stack)
}

func buildStack(tokenizer *xhtml.Tokenizer) []tagInfo {
	tokens := []tagInfo{}
	for tt := tokenizer.Next(); tt != xhtml.ErrorToken; tt = tokenizer.Next() {
		switch tt {
		case xhtml.TextToken:
			txt := string(tokenizer.Text())
			if len(tokens) == 0 {
				info := tagInfo{
					raw: txt,
				}
				tokens = append(tokens, info)
			}
			tn, _ := tokenizer.TagName()
			key, val, _ := tokenizer.TagAttr()
			info := tagInfo{
				tag: string(tn),
				key: string(key),
				val: string(val),
				txt: txt,
			}
			tokens = append(tokens, info)
		case xhtml.StartTagToken:
			tn, _ := tokenizer.TagName()
			key, val, _ := tokenizer.TagAttr()
			info := tagInfo{
				tag: string(tn),
				key: string(key),
				val: string(val),
			}
			tokens = append(tokens, info)
		case xhtml.SelfClosingTagToken, xhtml.EndTagToken:
			tn, _ := tokenizer.TagName()
			key, val, _ := tokenizer.TagAttr()
			info := tagInfo{
				tag:        string(tn),
				key:        string(key),
				val:        string(val),
				closingTag: true,
			}
			tokens = append(tokens, info)
		}
	}
	return tokens
}

func walk(tokens []tagInfo) string {
	// The first token can determine what type of tabbing
	doc := ""
	block := ""
	col := 0
	level := 0
	isIndented := false
	openTagValue := ""
	next := false

	for _, token := range tokens {
		if token.closingTag {
			level--
			if level == 0 {
				isIndented = false
			}
			block, col = wrap(block, col, 72, isIndented)
			endl := closeTag(token, level)
			block += endl
			if endl != "" {
				col = 0
			}
			doc += block
			block = ""
		} else {
			if token.raw != "" && token.raw != " " {
				doc += token.raw
				continue
			}
			// If the txt is blank, then it is a opening tag.
			// This is where indention will occur.
			if token.txt == "" {
				indentStr, indenting := indents(token.tag, level)
				isIndented = isIndented || indenting
				block += indentStr
				col += len(indentStr)
				openTagValue, next = formatText(token)
				if !next {
					block += openTagValue
				}
				level++
			} else {
				value, _ := formatText(token)
				block += value
				if next {
					block += openTagValue
				}
				next = false
			}
		}
	}
	return doc
}

// closeTag will divide up the blocks of documentation to be formated properly.
func closeTag(token tagInfo, level int) string {
	switch token.tag {
	case "pre", "li":
		return "\n"
	case "p", "h1", "h2", "h3", "h4", "h5", "h6", "div":
		// This is being inlined with something. We will ignore newlines
		if level > 0 {
			break
		}
		return "\n\n"
	case "code":
		if level == 0 {
			return "\n"
		}
	}
	return ""
}

func indents(tag string, level int) (string, bool) {
	switch tag {
	case "pre", "code":
		if level > 0 {
			break
		}
		fallthrough
	case "li":
		v, ok := style[tag]
		indentStr := indent
		if ok {
			indentStr += v
		}
		return indentStr, true
	}
	return "", false
}

// formatText will format any sort of text based off of a tag. It will also return
// a boolean to add the string after the text token.
func formatText(token tagInfo) (string, bool) {
	switch token.tag {
	case "a":
		if token.val != "" {
			return fmt.Sprintf(" (%s)", token.val), true
		}
	}
	if len(token.txt) == 0 || token.txt == " " {
		return "", false
	}
	return token.txt, false
}
