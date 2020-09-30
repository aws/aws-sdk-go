package api

import "text/template"

const endpointARNShapeTmplDef = `
{{- define "endpointARNShapeTmpl" }}
{{ range $_, $name := $.MemberNames -}}
	{{ $elem := index $.MemberRefs $name -}}
	{{ if $elem.EndpointARN -}}
		func (s *{{ $.ShapeName }}) getEndpointARN() (arn.Resource, error) {
			if s.{{ $name }} == nil {
				return nil, fmt.Errorf("member {{ $name }} is nil")
			}
			return parseEndpointARN(*s.{{ $name }})
		}

		func (s *{{ $.ShapeName }}) hasEndpointARN() bool {
			if s.{{ $name }} == nil {
				return false
			}
			return arn.IsARN(*s.{{ $name }})
		}

		// updateArnableField updates the value of the input field that 
		// takes an ARN as an input. This method is useful to backfill 
		// the parsed resource name from ARN into the input member. 
		func (s *{{ $.ShapeName }}) updateArnableField(v string) error {
			if s.{{ $name }} == nil {
				return fmt.Errorf("member {{ $name }} is nil")
			}
			s.{{ $name }} = aws.String(v)
			return nil
		}
	{{ end -}}
{{ end }}
{{ end }}
`

var endpointARNShapeTmpl = template.Must(
	template.New("endpointARNShapeTmpl").
		Parse(endpointARNShapeTmplDef),
)

const outpostIDShapeTmplDef = `
{{- define "outpostIDShapeTmpl" }}
{{ range $_, $name := $.MemberNames -}}
	{{ $elem := index $.MemberRefs $name -}}
	{{ if $elem.OutpostIDMember -}}
		func (s *{{ $.ShapeName }}) getOutpostID() (string, error) {
			if s.{{ $name }} == nil {
				return "", fmt.Errorf("member {{ $name }} is nil")
			}
			return *s.{{ $name }}, nil
		}

		func (s *{{ $.ShapeName }}) hasOutpostID() bool {
			if s.{{ $name }} == nil {
				return false
			}
			return true 
		}
	{{ end -}}
{{ end }}
{{ end }}
`

var outpostIDShapeTmpl = template.Must(
	template.New("outpostIDShapeTmpl").
		Parse(outpostIDShapeTmplDef),
)

const accountIDWithARNShapeTmplDef = `
{{- define "accountIDWithARNShapeTmpl" }}
{{ range $_, $name := $.MemberNames -}}
	{{ $elem := index $.MemberRefs $name -}}
	{{ if $elem.AccountIDMemberWithARN -}}
		func (s *{{ $.ShapeName }}) updateAccountID(accountId string) error {
			if s.{{ $name }} == nil {
				s.{{ $name }} = aws.String(accountId)	
			} else if *s.{{ $name }} != accountId  {
				return fmt.Errorf("Account ID mismatch, the Account ID cannot be specified in an ARN and in the accountId field")
			}
			return nil 
		}
	{{ end -}}
{{ end }}
{{ end }}
`

var accountIDWithARNShapeTmpl = template.Must(
	template.New("accountIDWithARNShapeTmpl").
		Parse(accountIDWithARNShapeTmplDef),
)
