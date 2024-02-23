# Domain Checker

## Overview
This Go program checks the domain-related information such as MX records, SPF records, and DMARC records for a list of specified domains. It provides information about whether the domain has MX records, SPF records, the SPF record itself, whether it has DMARC records, and the DMARC record itself.

## Features
- Checks domain-related information including MX, SPF, and DMARC records.
- Provides a comma-separated output format for easy analysis.

## Usage
1. Prepare a text file containing a list of domains, with each domain on a separate line.
2. Run the program by executing `go run main.go < input.txt`, where `input.txt` is the name of the file containing the list of domains.
3. The program will output the domain-related information in a comma-separated format, which can be redirected to a file or further processed as needed.

## Output Format
The program outputs the following fields in a comma-separated format:
- Domain: The domain being checked.
- HasMX: Indicates whether the domain has MX records (true/false).
- HasSPF: Indicates whether the domain has SPF records (true/false).
- SPFRecord: The SPF record associated with the domain (empty if not found).
- HasDMARC: Indicates whether the domain has DMARC records (true/false).
- DMARCRecord: The DMARC record associated with the domain (empty if not found).
