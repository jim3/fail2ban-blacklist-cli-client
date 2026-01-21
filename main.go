package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}
}

func main() {
	blacklist := flag.Int("blacklist", 0, "Number of banned IPs to look up (default: 1)")
	ip := flag.String("iplookup", "", "IP address to look up")
	cve := flag.String("cve", "", "Look up CVE vuln")
	flag.Parse()

	// Get list of banned IP's
	if *blacklist > 0 {
		var b BlacklistResponse
		b.GetBlacklist(*blacklist)
	}

	// Returns host information
	if *ip != "" {
		var resp IpLookUp
		err := resp.LookupIP(*ip)
		if err != nil {
			log.Fatalf("IP lookup failed: %v", err)
		}

		fmt.Println("===================================================")
		fmt.Printf("CPEs: %v\n", resp.CPES)
		fmt.Printf("Hostname: %v\n", resp.HostNames)
		fmt.Printf("IP: %v\n", resp.IP)
		fmt.Printf("Ports: %v\n", resp.Ports)
		fmt.Printf("Tags: %v\n", resp.Tags)
		fmt.Printf("Vulns: %v\n", resp.Vulns)
		fmt.Println()
		fmt.Println("===================================================")
	}

	// Look up CVE info
	if *cve != "" {
		var c CVE
		err := c.CveLookup(*cve)
		if err != nil {
			log.Fatalf("CVE lookup failed: %v", err)
		}

		fmt.Println("------------------------------- Summary -------------------------------")
		fmt.Println(c.Summary)
		fmt.Println("------------------------------- PublishedTime -------------------------------")
		fmt.Println(c.PublishedTime)
		fmt.Println("------------------------------- References -------------------------------")
		fmt.Println(c.References)
	}
}
