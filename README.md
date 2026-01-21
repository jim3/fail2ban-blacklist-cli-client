# Fail2Ban Blacklist CLI Client

A Go-based command-line tool for querying blacklisted IP addresses from my [Fail2Ban API server](https://github.com/jim3/fail2ban-blacklist-api). This client demonstrates practical security tooling by integrating threat intelligence from multiple sources.

## Features

- **Blacklist Retrieval**: Fetch IP addresses currently banned by Fail2Ban on my VPS
- **IP Intelligence**: Lookup detailed host information using Shodan's InternetDB
- **CVE Lookup**: Query vulnerability information from Shodan's CVE database

## Usage

```bash
# Get list of banned IPs and lookup details for the first 5 ([-count] option added soon)
go run . -blacklist # [-count 5]

# Look up specific IP address
go run . -iplookup "8.8.8.8"

# Look up CVE information
go run . -cve "CVE-2017-3646"
```

## Code Structure

### main.go
Entry point using Go's `flag` package for command-line argument parsing. Routes commands to appropriate handlers based on user flags.

### blacklist.go
Contains three main components:

1. **BlacklistResponse**: Fetches banned IPs from my Fail2Ban API at `https://<mydomain.com>/blacklist`
2. **IpLookUp**: Queries Shodan's InternetDB API for detailed host information (open ports, hostnames, CPEs, tags, vulnerabilities)
3. **CVE**: Retrieves vulnerability details from Shodan's CVE database

Each struct uses Go's `http` and `json` packages to handle REST API calls and response parsing.

## Dependencies

- Go standard library (`net/http`, `encoding/json`, `flag`)
- External APIs:
  - Personal Fail2Ban API
  - Shodan InternetDB (https://internetdb.shodan.io)
  - Shodan CVE Database (https://cvedb.shodan.io)

## Related Projects

This client pairs with my [Fail2Ban Blacklist API](https://github.com/jim3/fail2ban-blacklist-api) server, which exposes banned IPs from a live VPS.

## Status

This is a functional proof-of-concept demonstrating practical security tool development. While there are several TODOs and improvements planned (better error handling, configuration options, output formatting, etc.), it serves as a solid foundation for understanding API integration, threat intelligence gathering, and CLI tool design in Go.

Built as part of my portfolio for junior security analyst and penetration testing positions.


## Example Output

```bash
$ go run . -blacklist
===================================================
Looking up blacklisted ip address:  43.163.123.45
===================================================
cpes: [cpe:/a:openbsd:openssh:7.4]
Hostname: []
IP: 43.163.123.45
Open Ports: [22]
Tags: []
Vulns: [CVE-2020-14145 CVE-2025-32728 CVE-2021-36368 CVE-2023-51385 CVE-2018-20685 CVE-2019-6110 CVE-2019-6111 CVE-2023-48795 CVE-2018-15473 CVE-2008-3844 CVE-2023-38408 CVE-2020-15778 CVE-2016-20012 CVE-2007-2768 CVE-2017-15906 CVE-2023-51767 CVE-2019-6109 CVE-2025-26465 CVE-2018-15919 CVE-2021-41617]

===================================================
Looking up blacklisted ip address:  178.62.239.88
===================================================
cpes: [cpe:/a:ruby-lang:ruby cpe:/a:f5:nginx cpe:/a:rubyonrails:rails cpe:/a:openbsd:openssh:6.6.1p1 cpe:/o:canonical:ubuntu_linux]
Hostname: [gitlab.autorespond.nl]
IP: 178.62.239.88
Open Ports: [22 443]
Tags: [cloud self-signed]
Vulns: []

===================================================
Looking up blacklisted ip address:  181.64.27.145
===================================================
cpes: [cpe:/a:openbsd:openssh:9.2p1 cpe:/a:postfix:postfix cpe:/a:golang:go cpe:/a:apache:http_server:2.4.62 cpe:/o:linux:linux_kernel cpe:/a:apache:http_server cpe:/o:debian:debian_linux cpe:/a:caddyserver:caddy]
Hostname: [nube.santateresa.pe imagenes.santateresa.pe]
IP: 181.64.27.145
Open Ports: [22 25 80 443 465 587 2211 2222 3478 5000 8001 8006 8042 8069 8090 10000 20000]
Tags: [starttls]
Vulns: [CVE-2009-2299 CVE-2013-4365 CVE-2011-1176 CVE-2024-43204 CVE-2013-2765 CVE-2025-65082 CVE-2024-42516 CVE-2012-3526 CVE-2009-0796 CVE-2025-53020 CVE-2025-58098 CVE-2025-49630 CVE-2011-2688 CVE-2007-4723 CVE-2013-0942 CVE-2024-47252 CVE-2012-4001 CVE-2025-55753 CVE-2013-0941 CVE-2025-66200 CVE-2012-4360 CVE-2024-43394 CVE-2025-49812 CVE-2025-59775 CVE-2025-23048]

===================================================
```
