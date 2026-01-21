package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// For the JSON response, the Response struct
type BlacklistResponse struct {
	BannedIPs []string `json:"blacklist"`
}

func (b *BlacklistResponse) GetBlacklist() {
	var requestURL string = "https://<your-domain>/blacklist"
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client request failed: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client response failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("STATUS CODE:%d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	err = json.Unmarshal(resBody, &b)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	var count int = 5 // for testing, to keep the output "sane")
	for i := 0; i < count; i++ {
		var resp IpLookUp
		err = resp.LookupIP(b.BannedIPs[i])
		if err != nil {
			log.Fatalf("IP lookup failed: %v", err)
		}

		cpeStr := resp.CPES
		hostNames := resp.HostNames
		ipStr := resp.IP
		portsStr := resp.Ports
		tagsStr := resp.Tags
		vulnStr := resp.Vulns

		fmt.Println("Looking up blacklisted ip address: ", b.BannedIPs[i])
		fmt.Println("===================================================")
		fmt.Printf("cpes: %v\n", cpeStr)
		fmt.Printf("Hostname: %v\n", hostNames)
		fmt.Printf("IP: %v\n", ipStr)
		fmt.Printf("Open Ports: %v\n", portsStr)
		fmt.Printf("Tags: %v\n", tagsStr)
		fmt.Printf("Vulns: %v\n", vulnStr)
		fmt.Println()
		fmt.Println("===================================================")
	}

}

type IpLookUp struct {
	CPES      []string `json:"cpes"`
	HostNames []string `json:"hostnames"`
	IP        string   `json:"ip"`
	Ports     []int    `json:"ports"`
	Tags      []string `json:"tags"`
	Vulns     []string `json:"vulns"`
}

// Quick IP lookups via internetdb.shodan.io
func (i *IpLookUp) LookupIP(ipAddr string) error {
	URL := fmt.Sprintf("https://internetdb.shodan.io/%s", ipAddr)
	res, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading HostInfo body: %v", err)
	}

	err = json.Unmarshal(body, i)
	if err != nil {
		return fmt.Errorf("error unmarshalling json data: %v", err)
	}
	return nil
}

type CVE struct {
	Summary       string   `json:"summary"`
	PublishedTime string   `json:"published_time"`
	References    []string `json:"references"`
}

// Lookup CVE's
func (c *CVE) CveLookup(cve string) error {
	URL := fmt.Sprintf("https://cvedb.shodan.io/cve/%s", cve)
	res, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading HostInfo body: %v", err)
	}

	err = json.Unmarshal(body, c)
	if err != nil {
		return fmt.Errorf("error unmarshalling json data: %v", err)
	}
	return nil
}
