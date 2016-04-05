package main

import (
	"fmt"
	"github.com/jheise/gothreat"
)

func process_file(target string) {
	fmt.Printf("Looking up Hash: %s\n", target)
	file_data, err := gothreat.FileReport(target)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Permalink: %s\n", file_data.Permalink)
	fmt.Printf("References:\n")
	for _, reference := range file_data.References {
		fmt.Printf("\t%s\n", reference)
	}
	fmt.Printf("MD5: %s\n", file_data.Md5)
	fmt.Printf("SHA1: %s\n", file_data.Sha1)
	fmt.Printf("Scans:\n")
	for _, scan := range file_data.Scans {
		fmt.Printf("\t%s\n", scan)
	}
	fmt.Printf("IPs:\n")
	for _, ip := range file_data.Ips {
		fmt.Printf("\t%s\n", ip)
	}
	fmt.Printf("Domain:\n")
	for _, domain := range file_data.Domains {
		fmt.Printf("\t%s\n", domain)
	}
}
