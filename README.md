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

