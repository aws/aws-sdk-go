package api

import (
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// updateTopLevelResultWrappers moved resultWrappers from toplevel shape
// references to the toplevel shapes for easier code generation
func (a *API) updateTopLevelResultWrappers() {
	for _, o := range a.Operations {
		if o.InputRef.ResultWrapper != "" {
			o.InputRef.Shape.ResultWrapper = o.InputRef.ResultWrapper
		}
		if o.OutputRef.ResultWrapper != "" {
			o.OutputRef.Shape.ResultWrapper = o.OutputRef.ResultWrapper
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

func (a *API) renameToplevelShapes() {
	for _, v := range a.Operations {
		name := v.InputRef.ShapeName
		if strings.HasSuffix(name, "Request") {
			v.InputRef.Shape.Rename(strings.Replace(name, "Request", "Input", -1))
		}

		name = v.OutputRef.ShapeName
		if strings.HasSuffix(name, "Response") || strings.HasSuffix(name, "Result") {
			re := regexp.MustCompile(`(Response|Result)$`)
			v.OutputRef.Shape.Rename(re.ReplaceAllString(name, "Output"))
		}
	}
}

func (a *API) renameExportable() {
	for name, op := range a.Operations {
		newName := a.exportableName(name)
		if newName != name {
			delete(a.Operations, name)
			a.Operations[newName] = op
		}
		op.ExportedName = newName
	}

	for k, s := range a.Shapes {
		for mName, member := range s.MemberRefs {
			newName := a.exportableName(mName)
			if newName != mName {
				delete(s.MemberRefs, mName)
				s.MemberRefs[newName] = member

				// also apply locationName trait so we keep the old one
				if member.LocationName == "" {
					member.LocationName = mName
				}
			}

			if newName == "SDKShapeTraits" {
				panic("Shape " + s.ShapeName + " uses reserved member name SDKShapeTraits")
			}
		}

		newName := a.exportableName(k)
		if newName != s.ShapeName {
			s.Rename(newName)
		}

		s.Payload = a.exportableName(s.Payload)
		s.ResultWrapper = a.exportableName(s.ResultWrapper)
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

func (a *API) exportableName(name string) string {
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
	} else {
		return out
	}
}

var whitelistExportNames = func() map[string]string {
	list := map[string]string{}
	_, filename, _, _ := runtime.Caller(1)
	f, err := os.Open(path.Join(path.Dir(filename), "inflections.csv"))
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	str := string(b)
	for _, line := range strings.Split(str, "\n") {
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
