package main

import (
	"fmt"
	"github.com/jheise/gothreat"
)

func process_domain(target string) {

	fmt.Printf("Looking up Domain: %s\n", target)
	domain_data, err := gothreat.DomainReport(target)
	//body, err = gothreat.DomainReport(*target)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Permalink: %s\n", domain_data.Permalink)
	fmt.Printf("References:\n")
	for _, reference := range domain_data.References {
		fmt.Printf("\t%s\n", reference)
	}
	fmt.Printf("Hashes:\n")
	for _, hash := range domain_data.Hashes {
		fmt.Printf("\t%s\n", hash)
	}
	fmt.Printf("Emails:\n")
	for _, email := range domain_data.Emails {
		fmt.Printf("\t%s\n", email)
	}
	fmt.Printf("Subdomains:\n")
	for _, subdomain := range domain_data.Subdomains {
		fmt.Printf("\t%s\n", subdomain)
	}
	fmt.Printf("Resolutions:\n")
	for _, resolve := range domain_data.Resolutions {
		fmt.Printf("\t%s -> %s\n", resolve.LastResolved, resolve.IPAddress)
	}
}
