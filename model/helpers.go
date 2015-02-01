package model

import (
	"bytes"
	"go/doc"
	"regexp"
	"sort"
	"strings"

	"code.google.com/p/go.net/html"
	"github.com/aarzilli/sandblast"
)

func locationTraits(location string, shape *Shape) string {
	if shape == nil {
		return ""
	}

	var list []string
	for name, member := range shape.Members() {
		if member.Location == location {
			list = append(list, "\""+name+"\"")
		}
	}
	sort.Strings(list)
	return strings.Join(list, ",")
}

func requiredTraits(shape *Shape) string {
	if shape == nil {
		return ""
	}

	var list []string
	for name, member := range shape.Members() {
		if member.Required {
			list = append(list, "\""+name+"\"")
		}
	}
	sort.Strings(list)
	return strings.Join(list, ",")
}

func hasRequiredTrait(shape *Shape) bool {
	if shape != nil {
		for _, member := range shape.Members() {
			if member.Required {
				return true
			}
		}
	}
	return false
}

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

func protocolPackage(protocol string) string {
	switch {
	case protocol == "json":
		return "jsonrpc"
	default:
		return strings.Replace(protocol, "-", "", -1)
	}
}

func structName(name string) string {
	str := exportable(name)
	str = regexp.MustCompile(`Request$`).ReplaceAllString(str, "Input")
	str = regexp.MustCompile(`Response$`).ReplaceAllString(str, "Output")
	str = regexp.MustCompile(`Result$`).ReplaceAllString(str, "Output")
	return str
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

type shapeMapEntry struct {
	*Shape
	Alias    string
	TopLevel bool
}

type shapeMapEntryList []shapeMapEntry

func (l shapeMapEntryList) Len() int {
	return len(l)
}

func (l shapeMapEntryList) Less(i, j int) bool {
	s := sort.StringSlice{l[i].Name, l[j].Name}
	return s.Less(0, 1)
}

func (l shapeMapEntryList) Swap(i, j int) {
	t := l[i]
	l[i] = l[j]
	l[j] = t
}

var shapeMap = map[string]shapeMapEntry{}
var enumMap = map[string]shapeMapEntry{}

func buildShapeMap() {
	if len(shapeMap) == 0 {
		for _, v := range service.Operations {
			buildShapeMapForShape(v.Input(), true)
			buildShapeMapForShape(v.Output(), true)
		}
	}
}

func buildShapeMapForShape(shape *Shape, toplevel bool) {
	if shape == nil || shape.Exception {
		return
	}

	if shape.Enum != nil {
		if _, exists := enumMap[shape.Name]; !exists {
			enumMap[shape.Name] = shapeMapEntry{
				Shape: shape,
				Alias: exportable(shape.Name),
			}
		}
		return
	}

	if _, exists := shapeMap[shape.Name]; exists {
		return
	}

	if toplevel {
		shapeMap[shape.Name] = shapeMapEntry{
			Shape:    shape,
			TopLevel: toplevel,
			Alias:    structName(shape.Name),
		}
	} else {
		shapeMap[shape.Name] = shapeMapEntry{
			Shape:    shape,
			TopLevel: toplevel,
			Alias:    exportable(shape.Name),
		}
	}

	switch shape.ShapeType {
	case "structure":
		for _, member := range shape.Members() {
			buildShapeMapForShape(member.Shape(), false)
		}

	case "map":
		buildShapeMapForShape(shape.Key(), false)
		buildShapeMapForShape(shape.Value(), false)

	case "list":
		buildShapeMapForShape(shape.Member(), false)

	}
}

func shapeList() shapeMapEntryList {
	list := make(shapeMapEntryList, len(shapeMap))

	i := 0
	for _, v := range shapeMap {
		list[i] = v
		i++
	}
	sort.Sort(list)

	return list
}

func enumList() shapeMapEntryList {
	list := make(shapeMapEntryList, len(enumMap))

	i := 0
	for _, v := range enumMap {
		list[i] = v
		i++
	}
	sort.Sort(list)

	return list
}

func shapeAlias(name string) string {
	return shapeMap[name].Alias
}

var replacements = map[*regexp.Regexp]string{
	regexp.MustCompile(`Acl`):          "ACL",
	regexp.MustCompile(`Adm([^i]|$)`):  "ADM$1",
	regexp.MustCompile(`Aes`):          "AES",
	regexp.MustCompile(`Api`):          "API",
	regexp.MustCompile(`Ami`):          "AMI",
	regexp.MustCompile(`Apns`):         "APNS",
	regexp.MustCompile(`Arn`):          "ARN",
	regexp.MustCompile(`Asn`):          "ASN",
	regexp.MustCompile(`Aws`):          "AWS",
	regexp.MustCompile(`Bcc([A-Z])`):   "BCC$1",
	regexp.MustCompile(`Bgp`):          "BGP",
	regexp.MustCompile(`Cc([A-Z])`):    "CC$1",
	regexp.MustCompile(`Cidr`):         "CIDR",
	regexp.MustCompile(`Cors`):         "CORS",
	regexp.MustCompile(`Csv`):          "CSV",
	regexp.MustCompile(`Cpu`):          "CPU",
	regexp.MustCompile(`Db`):           "DB",
	regexp.MustCompile(`Dhcp`):         "DHCP",
	regexp.MustCompile(`Dns`):          "DNS",
	regexp.MustCompile(`Ebs`):          "EBS",
	regexp.MustCompile(`Ec2`):          "EC2",
	regexp.MustCompile(`Eip`):          "EIP",
	regexp.MustCompile(`Gcm`):          "GCM",
	regexp.MustCompile(`Html`):         "HTML",
	regexp.MustCompile(`Https`):        "HTTPS",
	regexp.MustCompile(`Http([^s]|$)`): "HTTP$1",
	regexp.MustCompile(`Hsm`):          "HSM",
	regexp.MustCompile(`Hvm`):          "HVM",
	regexp.MustCompile(`Iam`):          "IAM",
	regexp.MustCompile(`Icmp`):         "ICMP",
	regexp.MustCompile(`Id$`):          "ID",
	regexp.MustCompile(`Id([A-Z])`):    "ID$1",
	regexp.MustCompile(`Idn`):          "IDN",
	regexp.MustCompile(`Ids$`):         "IDs",
	regexp.MustCompile(`Ids([A-Z])`):   "IDs$1",
	regexp.MustCompile(`Iops`):         "IOPS",
	regexp.MustCompile(`Ip`):           "IP",
	regexp.MustCompile(`Jar`):          "JAR",
	regexp.MustCompile(`Json`):         "JSON",
	regexp.MustCompile(`Jvm`):          "JVM",
	regexp.MustCompile(`Kms`):          "KMS",
	regexp.MustCompile(`Mac([^h]|$)`):  "MAC$1",
	regexp.MustCompile(`Md5`):          "MD5",
	regexp.MustCompile(`Mfa`):          "MFA",
	regexp.MustCompile(`Ok`):           "OK",
	regexp.MustCompile(`Os`):           "OS",
	regexp.MustCompile(`Php`):          "PHP",
	regexp.MustCompile(`Raid`):         "RAID",
	regexp.MustCompile(`Ramdisk`):      "RAMDisk",
	regexp.MustCompile(`Rds`):          "RDS",
	regexp.MustCompile(`Sni`):          "SNI",
	regexp.MustCompile(`Sns`):          "SNS",
	regexp.MustCompile(`Sriov`):        "SRIOV",
	regexp.MustCompile(`Ssh`):          "SSH",
	regexp.MustCompile(`Ssl`):          "SSL",
	regexp.MustCompile(`Svn`):          "SVN",
	regexp.MustCompile(`Tar([^g]|$)`):  "TAR$1",
	regexp.MustCompile(`Tde`):          "TDE",
	regexp.MustCompile(`Tcp`):          "TCP",
	regexp.MustCompile(`Tgz`):          "TGZ",
	regexp.MustCompile(`Tls`):          "TLS",
	regexp.MustCompile(`Uri`):          "URI",
	regexp.MustCompile(`Url`):          "URL",
	regexp.MustCompile(`Vgw`):          "VGW",
	regexp.MustCompile(`Vhd`):          "VHD",
	regexp.MustCompile(`Vip`):          "VIP",
	regexp.MustCompile(`Vlan`):         "VLAN",
	regexp.MustCompile(`Vm([^d]|$)`):   "VM$1",
	regexp.MustCompile(`Vmdk`):         "VMDK",
	regexp.MustCompile(`Vpc`):          "VPC",
	regexp.MustCompile(`Vpn`):          "VPN",
	regexp.MustCompile(`Xml`):          "XML",
}
