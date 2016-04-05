package main

import (
	"fmt"
	"github.com/jheise/gothreat"
)

func process_av(target string) {
	fmt.Printf("Looking up Antivirus: %s\n", target)
	av_data, err := gothreat.AntiVirusReport(target)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Permalink: %s\n", av_data.Permalink)
	fmt.Printf("References:\n")
	for _, reference := range av_data.References {
		fmt.Printf("\t%s\n", reference)
	}
	fmt.Printf("Hashes:\n")
	for _, hash := range av_data.Hashes {
		fmt.Printf("\t%s\n", hash)
	}
}
