package api

var svcCustomizations = map[string]func(*API){
	"s3": s3Customizations,
}

func (a *API) customizationPasses() {
	if fn := svcCustomizations[a.PackageName()]; fn != nil {
		fn(a)
	}
}

func s3Customizations(a *API) {
	// Remove ContentMD5 members
	for _, s := range a.Shapes {
		if _, ok := s.MemberRefs["ContentMD5"]; ok {
			delete(s.MemberRefs, "ContentMD5")
		}
	}

	// Rename "Rule" to "LifecycleRule"
	if s, ok := a.Shapes["Rule"]; ok {
		s.Rename("LifecycleRule")
	}
}
