package api

import (
	"fmt"
	"regexp"
	"strings"
)

func (a *API) writeShapeNames() {
	for n, s := range a.Shapes {
		s.API = a
		s.ShapeName = n
	}
}

func (a *API) resolveReferences() {
	resolver := referenceResolver{API: a, visited: map[*ShapeRef]bool{}}

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
		fmt.Println("Adding ref", ref.ShapeName, "to", shape.ShapeName)
		ref.API = r.API                      // resolve reference back to API
		ref.Shape = shape                    // resolve shape reference
		shape.refs = append(shape.refs, ref) // register the ref

		if r.visited[ref] {
			return
		}
		r.visited[ref] = true

		// resolve shape's references, if it has any
		r.resolveReference(&shape.MemberRef)
		r.resolveReference(&shape.KeyRef)
		r.resolveReference(&shape.ValueRef)
		for _, m := range shape.MemberRefs {
			r.resolveReference(&m)
		}
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
		newName := exportableName(name)
		if newName != name {
			delete(a.Operations, name)
			a.Operations[newName] = op
		}
		op.ExportedName = newName
	}

	for k, s := range a.Shapes {
		for mName, member := range s.MemberRefs {
			newName := exportableName(mName)
			if newName != mName {
				delete(s.MemberRefs, mName)
				s.MemberRefs[newName] = member
			}
		}

		newName := exportableName(k)
		if newName != s.ShapeName {
			s.Rename(newName)
		}
	}
}

func exportableName(name string) string {
	// make sure the symbol is exportable
	name = strings.ToUpper(name[0:1]) + name[1:]

	// fix common AWS<->Go bugaboos
	for regexp, repl := range replacements {
		name = regexp.ReplaceAllString(name, repl)
	}
	return name
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
