package main

import (
	"fmt"
	"github.com/jheise/gothreat"
)

func process_ip(target string) {
	fmt.Printf("Looking up IP: %s\n", target)
	ip_data, err := gothreat.IPReport(target)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Permalink: %s\n", ip_data.Permalink)
	fmt.Printf("DNS Resolutions:\n")
	for _, resolve := range ip_data.Resolutions {
		fmt.Printf("\t%s -> %s\n", resolve.LastResolved, resolve.Domain)
	}
	fmt.Printf("References:\n")
	for _, reference := range ip_data.References {
		fmt.Printf("\t%s\n", reference)
	}
	fmt.Printf("Hashes:\n")
	for _, hash := range ip_data.Hashes {
		fmt.Printf("\t%s\n", hash)
	}
}
