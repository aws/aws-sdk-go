package aws

//go:generate go run ../internal/model/cli/gen-endpoints/main.go endpoints.json endpoints_map.go

import "strings"

func (s Service) endpointForRegion() string {
	derivedKeys := []string{
		s.Config.Region + "/" + s.ServiceName,
		s.Config.Region + "/*",
		"*/" + s.ServiceName,
		"*/*",
	}

	for _, key := range derivedKeys {
		if val, ok := endpointsMap.Endpoints[key]; ok {
			ep := val.Endpoint
			ep = strings.Replace(ep, "{region}", s.Config.Region, -1)
			ep = strings.Replace(ep, "{service}", s.ServiceName, -1)
			return ep
		}
	}
	return ""
}
