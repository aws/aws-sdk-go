package api

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// updateTopLevelShapeReferences moves resultWrapper, locationName, and
// xmlNamespace traits from toplevel shape references to the toplevel
// shapes for easier code generation
func (a *API) updateTopLevelShapeReferences() {
	for _, o := range a.Operations {
		// these are for Query services
		if o.InputRef.ResultWrapper != "" {
			o.InputRef.Shape.ResultWrapper = o.InputRef.ResultWrapper
		}
		if o.OutputRef.ResultWrapper != "" {
			o.OutputRef.Shape.ResultWrapper = o.OutputRef.ResultWrapper
		}

		// these are for REST-XML services
		if o.InputRef.LocationName != "" {
			o.InputRef.Shape.LocationName = o.InputRef.LocationName
		}
		if o.InputRef.Location != "" {
			o.InputRef.Shape.Location = o.InputRef.Location
		}
		if o.InputRef.Payload != "" {
			o.InputRef.Shape.Payload = o.InputRef.Payload
		}
		if o.InputRef.XMLNamespace.Prefix != "" {
			o.InputRef.Shape.XMLNamespace.Prefix = o.InputRef.XMLNamespace.Prefix
		}
		if o.InputRef.XMLNamespace.URI != "" {
			o.InputRef.Shape.XMLNamespace.URI = o.InputRef.XMLNamespace.URI
		}
	}

}

func (a *API) writeShapeNames() {
	for n, s := range a.Shapes {
		s.API = a
		s.ShapeName = n
	}
}

func (a *API) resolveReferences() {
	resolver := referenceResolver{API: a, visited: map[*ShapeRef]bool{}}

	for _, s := range a.Shapes {
		resolver.resolveShape(s)
	}

	for _, o := range a.Operations {
		o.API = a // resolve parent reference

		resolver.resolveReference(&o.InputRef)
		resolver.resolveReference(&o.OutputRef)
	}
}

type referenceResolver struct {
	*API
	visited map[*ShapeRef]bool
}

func (r *referenceResolver) resolveReference(ref *ShapeRef) {
	if ref.ShapeName == "" {
		return
	}

	if shape, ok := r.API.Shapes[ref.ShapeName]; ok {
		ref.API = r.API   // resolve reference back to API
		ref.Shape = shape // resolve shape reference

		if r.visited[ref] {
			return
		}
		r.visited[ref] = true

		shape.refs = append(shape.refs, ref) // register the ref

		// resolve shape's references, if it has any
		r.resolveShape(shape)
	}
}

func (r *referenceResolver) resolveShape(shape *Shape) {
	r.resolveReference(&shape.MemberRef)
	r.resolveReference(&shape.KeyRef)
	r.resolveReference(&shape.ValueRef)
	for _, m := range shape.MemberRefs {
		r.resolveReference(m)
	}
}

func removeReference(refs []*ShapeRef, ref *ShapeRef) []*ShapeRef {
	list := []*ShapeRef{}
	for _, v := range refs {
		if v != ref {
			list = append(list, ref)
		}
	}
	return list
}

func (a *API) createShapeFromRef(ref *ShapeRef, name string) {
	// create a new copy of this shape
	copy := *ref.Shape
	copy.ShapeName = name
	copy.refs = []*ShapeRef{ref}
	removeReference(ref.Shape.refs, ref)

	// add this copy to shape list and update ref
	a.Shapes[name] = &copy
	ref.ShapeName = name
}

func (a *API) renameToplevelShapes() {
	for _, v := range a.Operations {
		if v.HasInput() {
			name := v.ExportedName + "Input"
			switch n := len(v.InputRef.Shape.refs); {
			case n == 1:
				v.InputRef.Shape.Rename(name)
			case n > 1 && v.InputRef.ResultWrapper != "":
				a.createShapeFromRef(&v.InputRef, name)
			}
		}
		if v.HasOutput() {
			name := v.ExportedName + "Output"
			switch n := len(v.OutputRef.Shape.refs); {
			case n == 1:
				v.OutputRef.Shape.Rename(name)
			case n > 1 && v.OutputRef.ResultWrapper != "":
				a.createShapeFromRef(&v.OutputRef, name)
			}
		}
		v.InputRef.Payload = a.ExportableName(v.InputRef.Payload)
		v.OutputRef.Payload = a.ExportableName(v.OutputRef.Payload)
	}
}

func (a *API) renameExportable() {
	for name, op := range a.Operations {
		newName := a.ExportableName(name)
		if newName != name {
			delete(a.Operations, name)
			a.Operations[newName] = op
		}
		op.ExportedName = newName
	}

	for k, s := range a.Shapes {
		// FIXME SNS has lower and uppercased shape names with the same name,
		// except the lowercased variant is used exclusively for string and
		// other primitive types. Renaming both would cause a collision.
		// We work around this by only renaming the structure shapes.
		if s.Type == "string" {
			continue
		}

		for mName, member := range s.MemberRefs {
			newName := a.ExportableName(mName)
			if newName != mName {
				delete(s.MemberRefs, mName)
				s.MemberRefs[newName] = member

				// also apply locationName trait so we keep the old one
				// but only if there's no locationName trait on ref or shape
				if member.LocationName == "" && member.Shape.LocationName == "" {
					member.LocationName = mName
				}
			}

			if newName == "SDKShapeTraits" {
				panic("Shape " + s.ShapeName + " uses reserved member name SDKShapeTraits")
			}
		}

		newName := a.ExportableName(k)
		if newName != s.ShapeName {
			s.Rename(newName)
		}

		s.Payload = a.ExportableName(s.Payload)
		s.ResultWrapper = a.ExportableName(s.ResultWrapper)
	}
}

// createInputOutputShapes creates toplevel input/output shapes if they
// have not been defined in the API. This normalizes all APIs to always
// have an input and output structure in the signature.
func (a *API) createInputOutputShapes() {
	for _, v := range a.Operations {
		if !v.HasInput() {
			shape := a.makeIOShape(v.ExportedName + "Input")
			v.InputRef = ShapeRef{API: a, ShapeName: shape.ShapeName, Shape: shape}
			shape.refs = append(shape.refs, &v.InputRef)
		}
		if !v.HasOutput() {
			shape := a.makeIOShape(v.ExportedName + "Output")
			v.OutputRef = ShapeRef{API: a, ShapeName: shape.ShapeName, Shape: shape}
			shape.refs = append(shape.refs, &v.OutputRef)
		}
	}
}

func (a *API) makeIOShape(name string) *Shape {
	shape := &Shape{
		API: a, ShapeName: name, Type: "structure",
		MemberRefs: map[string]*ShapeRef{},
	}
	a.Shapes[name] = shape
	return shape
}

func (a *API) removeUnusedShapes() {
	for n, s := range a.Shapes {
		if len(s.refs) == 0 {
			delete(a.Shapes, n)
		}
	}
}

func splitName(name string) []string {
	out, buf := []string{}, ""

	for i, r := range name {
		l := string(r)

		// special check for EC2 or MD5
		if _, err := strconv.Atoi(l); err == nil && (strings.ToLower(buf) == "ec" || strings.ToLower(buf) == "md") {
			buf += l
			continue
		}

		lastUpper := i-1 >= 0 && strings.ToUpper(name[i-1:i]) == name[i-1:i]
		curUpper := l == strings.ToUpper(l)
		nextUpper := i+2 > len(name) || strings.ToUpper(name[i+1:i+2]) == name[i+1:i+2]

		if (lastUpper != curUpper) || (nextUpper != curUpper && !nextUpper) {
			if len(buf) > 1 || curUpper {
				out = append(out, buf)
				buf = ""
			}
			buf += l
		} else {
			buf += l
		}
	}
	if len(buf) > 0 {
		out = append(out, buf)
	}
	return out
}

func (a *API) ExportableName(name string) string {
	if name == "" {
		return name
	}

	failed := false

	// make sure the symbol is exportable
	name = strings.ToUpper(name[0:1]) + name[1:]

	// inflections are disabled, stop here.
	if a.NoInflections {
		return name
	}

	// fix common AWS<->Go bugaboos
	out := ""
	for _, part := range splitName(name) {
		if part == "" {
			continue
		}
		if part == strings.ToUpper(part) || part[0:1]+"s" == part {
			out += part
			continue
		}
		if v, ok := whitelistExportNames[part]; ok {
			if v != "" {
				out += v
			} else {
				out += part
			}
		} else {
			failed = true
			inflected := part
			for regexp, repl := range replacements {
				inflected = regexp.ReplaceAllString(inflected, repl)
			}
			a.unrecognizedNames[part] = inflected
		}
	}

	if failed {
		return name
	}
	return out
}

var whitelistExportNames = func() map[string]string {
	list := map[string]string{}
	_, filename, _, _ := runtime.Caller(1)
	f, err := os.Open(filepath.Join(filepath.Dir(filename), "inflections.csv"))
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	str := string(b)

	if f, err := os.Open("inflections.csv"); err == nil {
		if additionalInflections, err := ioutil.ReadAll(f); err == nil {
			str += "\n" + string(additionalInflections)
		}
	}

	for _, line := range strings.Split(str, "\n") {
		line = strings.Replace(line, "\r", "", -1)
		if strings.HasPrefix(line, ";") {
			continue
		}
		parts := regexp.MustCompile(`\s*:\s*`).Split(line, -1)
		if len(parts) > 1 {
			list[parts[0]] = parts[1]
		}
	}

	return list
}()

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
