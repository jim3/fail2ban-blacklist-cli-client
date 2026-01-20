package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	blacklist := flag.Bool("blacklist", false, "Get banned IP's")
	ip := flag.String("iplookup", "", "IP address to look up")
	cve := flag.String("cve", "", "Look up CVE vuln")
	flag.Parse()

	// Get list of banned IP's
	if *blacklist {
		var b BlacklistResponse
		b.GetBlacklist()
	}

	// Returns host information
	if *ip != "" {
		var resp IpLookUp
		err := resp.LookupIP(*ip)
		if err != nil {
			log.Fatalf("IP lookup failed: %v", err)
		}

		cpeStr := resp.CPES
		hostNames := resp.HostNames
		ipStr := resp.IP
		portsStr := resp.Ports
		tagsStr := resp.Tags
		vulnStr := resp.Vulns
		fmt.Println("===================================================")
		fmt.Printf("cpes: %v\n", cpeStr)
		fmt.Printf("hostnames: %v\n", hostNames)
		fmt.Printf("ip: %v\n", ipStr)
		fmt.Printf("ports: %v\n", portsStr)
		fmt.Printf("tags: %v\n", tagsStr)
		fmt.Printf("vulns: %v\n", vulnStr)
		fmt.Println("===================================================")
	}

	// Look up CVE info
	if *cve != "" {
		var c CVE
		err := c.CveLookup(*cve)
		if err != nil {
			log.Fatalf("CVE lookup failed: %v", err)
		}
		fmt.Println(c.Summary)
		fmt.Println(c.PublishedTime)
		fmt.Println(c.References)
	}
}
