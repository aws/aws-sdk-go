package api

// updateTopLevelShapeReferences moves resultWrapper, locationName, and
// xmlNamespace traits from toplevel shape references to the toplevel
// shapes for easier code generation
func (a *API) updateTopLevelShapeReferences() {
	for _, o := range a.Operations {
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

func (a *API) renameToplevelShapes() {
	for _, v := range a.Operations {
		if v.HasInput() && !a.NoInflections {
			name := v.ExportedName + "Input"
			switch n := len(v.InputRef.Shape.refs); {
			case n == 1:
				v.InputRef.Shape.Rename(name)
			}
		}
		if v.HasOutput() && !a.NoInflections {
			name := v.ExportedName + "Output"
			switch n := len(v.OutputRef.Shape.refs); {
			case n == 1:
				v.OutputRef.Shape.Rename(name)
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

		// fix required trait names
		for i, n := range s.Required {
			s.Required[i] = a.ExportableName(n)
		}
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
