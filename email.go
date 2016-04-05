package main

import (
	"fmt"
	"github.com/jheise/gothreat"
)

func process_email(target string) {
	fmt.Printf("Looking up Email: %s\n", target)
	email_data, err := gothreat.EmailReport(target)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Permalink: %s\n", email_data.Permalink)
	fmt.Printf("References:\n")
	for _, reference := range email_data.References {
		fmt.Printf("\t%s\n", reference)
	}
	fmt.Printf("Domains:\n")
	for _, domains := range email_data.Domains {
		fmt.Printf("\t%s\n", domains)
	}
}
