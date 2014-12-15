package model

import (
	"bytes"
	"go/doc"
	"regexp"
	"strings"

	"code.google.com/p/go.net/html"
	"github.com/aarzilli/sandblast"
)

func godoc(member, content string) string {
	undocumented := "// " + exportable(member) + " is undocumented.\n"

	node, err := html.Parse(strings.NewReader(content))
	if err != nil {
		return undocumented
	}

	_, v, err := sandblast.Extract(node)
	if err != nil {
		return undocumented
	}

	v = strings.TrimSpace(v)
	if v == "" {
		return undocumented
	}

	if member != "" {
		v = exportable(member) + " " + strings.ToLower(v[0:1]) + v[1:]
	}

	out := bytes.NewBuffer(nil)
	doc.ToText(out, v, "// ", "", 72)
	return out.String()
}

func exportable(name string) string {
	// make sure the symbol is exportable
	name = strings.ToUpper(name[0:1]) + name[1:]

	// fix common AWS<->Go bugaboos
	for regexp, repl := range replacements {
		name = regexp.ReplaceAllString(name, repl)
	}
	return name
}

var replacements = map[*regexp.Regexp]string{
	regexp.MustCompile(`Api`):        "API",
	regexp.MustCompile(`Arn`):        "ARN",
	regexp.MustCompile(`Asn`):        "ASN",
	regexp.MustCompile(`Bgp`):        "BGP",
	regexp.MustCompile(`Cidr`):       "CIDR",
	regexp.MustCompile(`Cpu`):        "CPU",
	regexp.MustCompile(`Dhcp`):       "DHCP",
	regexp.MustCompile(`Dns`):        "DNS",
	regexp.MustCompile(`Ebs`):        "EBS",
	regexp.MustCompile(`Ec2`):        "EC2",
	regexp.MustCompile(`Html`):       "HTML",
	regexp.MustCompile(`Http`):       "HTTP",
	regexp.MustCompile(`Id$`):        "ID",
	regexp.MustCompile(`Id([A-Z])`):  "ID$1",
	regexp.MustCompile(`Ids$`):       "IDs",
	regexp.MustCompile(`Ids([A-Z])`): "IDs$1",
	regexp.MustCompile(`Ip`):         "IP",
	regexp.MustCompile(`Json`):       "JSON",
	regexp.MustCompile(`Sns`):        "SNS",
	regexp.MustCompile(`Ssh`):        "SSH",
	regexp.MustCompile(`Uri`):        "URI",
	regexp.MustCompile(`Url`):        "URL",
	regexp.MustCompile(`Vlan`):       "VLAN",
	regexp.MustCompile(`Vpc`):        "VPC",
}
