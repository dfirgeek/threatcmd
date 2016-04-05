package main

import (
	"github.com/alecthomas/kingpin"
)

var (
	//target string
	report = kingpin.Arg("report", "What to report on [av, domain email, file, ip]").Required().String()
	target = kingpin.Arg("target", "What to report on").Required().String()
)

func main() {
	kingpin.Parse()

	if *report == "ip" {
		process_ip(*target)
	} else if *report == "domain" {
		process_domain(*target)
	} else if *report == "email" {
		process_email(*target)
	} else if *report == "av" {
		process_av(*target)
	} else if *report == "file" {
		process_file(*target)
	}

}
