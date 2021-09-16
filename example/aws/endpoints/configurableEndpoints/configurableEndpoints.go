//go:build example
// +build example

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/endpoints"
)


// Command: go run enumEndpoints.go -partitions -regions -services -p minio -r us-east-1
func main() {
	var partitionID, regionID, serviceID string
	flag.StringVar(&partitionID, "p", "", "Partition ID")
	flag.StringVar(&regionID, "r", "", "Region ID")
	flag.StringVar(&serviceID, "s", "", "Service ID")

	var cmdPartitions, cmdRegions, cmdServices bool
	flag.BoolVar(&cmdPartitions, "partitions", false, "Lists partitions.")
	flag.BoolVar(&cmdRegions, "regions", false, "Lists regions of a partition. Requires partition ID to be provided. Will filter by a service if '-s' is set.")
	flag.BoolVar(&cmdServices, "services", false, "Lists services for a partition. Requires partition ID to be provided. Will filter by a region if '-r' is set.")
	flag.Parse()

	os.Setenv("ENDPOINT_CONFIG_FILE", "./endpointConfig.json")

	partitions := endpoints.DefaultResolver().(endpoints.EnumPartitions).Partitions()

	if cmdPartitions {
		printPartitions(partitions)
	}

	if !(cmdRegions || cmdServices) {
		return
	}

	p, ok := findPartition(partitions, partitionID)
	if !ok {
		fmt.Fprintf(os.Stderr, "Partition %q not found", partitionID)
		os.Exit(1)
	}
	fmt.Println()

	if cmdRegions {
		printRegions(p, serviceID)
	}
	fmt.Println()

	if cmdServices {
		printServices(p, regionID)
	}
	fmt.Println()
}

func printPartitions(ps []endpoints.Partition) {
	fmt.Println("Partitions:")
	for _, p := range ps {
		fmt.Println(p.ID())
	}
}

func printRegions(p endpoints.Partition, serviceID string) {
	if len(serviceID) != 0 {
		s, ok := p.Services()[serviceID]
		if !ok {
			fmt.Fprintf(os.Stderr, "service %q does not exist in partition %q", serviceID, p.ID())
			os.Exit(1)
		}
		es := s.Endpoints()
		fmt.Printf("Endpoints for %s in %s:\n", serviceID, p.ID())
		for _, e := range es {
			r, _ := e.ResolveEndpoint()
			fmt.Printf("%s: %s\n", e.ID(), r.URL)
		}

	} else {
		rs := p.Regions()
		fmt.Printf("Regions in %s:\n", p.ID())
		for _, r := range rs {
			fmt.Println(r.ID())
		}
	}
}

func printServices(p endpoints.Partition, endpointID string) {
	ss := p.Services()

	if len(endpointID) > 0 {
		fmt.Printf("Services with endpoint %s in %s:\n", endpointID, p.ID())
	} else {
		fmt.Printf("Services in %s:\n", p.ID())
	}

	for id, s := range ss {
		if _, ok := s.Endpoints()[endpointID]; !ok && len(endpointID) > 0 {
			continue
		}
		fmt.Println(id)
	}
}

func findPartition(ps []endpoints.Partition, partitionID string) (endpoints.Partition, bool) {
	for _, p := range ps {
		if p.ID() == partitionID {
			return p, true
		}
	}

	return endpoints.Partition{}, false
}
