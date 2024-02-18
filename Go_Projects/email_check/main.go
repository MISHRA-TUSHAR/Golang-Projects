package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord")

	var wg sync.WaitGroup
	for scanner.Scan() {
		domain := scanner.Text()
		wg.Add(1)
		go func(domain string) {
			defer wg.Done()
			checkDomain(domain)
		}(domain)
	}
	wg.Wait()

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read the input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: could not find MX records for %s: %v\n", domain, err)
		return
	}

	hasMX := len(mxRecords) > 0

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: could not find TXT records for %s: %v\n", domain, err)
		return
	}

	for _, record := range txtRecords {
		if strings.Contains(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
		}

		if strings.Contains(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
		}
	}

	fmt.Printf("%s,%t,%t,%s,%t,%s\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
