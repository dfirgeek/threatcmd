package main

import (
    "fmt"
    "github.com/jheise/gothreat"
    "github.com/alecthomas/kingpin"
)

var (
    //target string
    report = kingpin.Arg("report", "What to report on").Required().String()
    target = kingpin.Arg("target", "What to report on").Required().String()

)

func main(){
    kingpin.Parse()

    if *report == "ip" {
        fmt.Printf("Looking up IP: %s\n", *target)
        ip_data,err := gothreat.IPReport(*target)

        if err != nil {
            panic(err)
        }

        fmt.Printf("Permalink: %s\n", ip_data.Permalink)
        fmt.Printf("DNS Resolutions:\n")
        for _, resolve := range ip_data.Resolutions {
            fmt.Printf("\t%s -> %s\n", resolve["last_resolved"], resolve["domain"])
        }
        fmt.Printf("References:\n")
        for _, reference := range ip_data.References {
            fmt.Printf("\t%s\n", reference)
        }
        fmt.Printf("Hashes:\n")
        for _, hash := range ip_data.Hashes {
            fmt.Printf("\t%s\n", hash)
        }
    } else if *report == "domain" {
        fmt.Printf("Looking up Domain: %s\n", *target)
        domain_data, err := gothreat.DomainReport(*target)
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
            fmt.Printf("\t%s -> %s\n", resolve["last_resolved"], resolve["ip_address"])
        }

    } else if *report == "email" {
        fmt.Printf("Looking up Email: %s\n", *target)
        email_data, err := gothreat.EmailReport(*target)
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
    } else if *report == "av" {
        fmt.Printf("Looking up Antivirus: %s\n", *target)
        av_data, err := gothreat.AntiVirusReport(*target)
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


}
