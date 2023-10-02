package main

import (
	"bufio"     // Importing the bufio package to read from standard input
	"fmt"      // Importing the fmt package for formatting and printing
	"log"      // Importing the log package to log errors
	"net"      // Importing the net package for network I/O functions
	"os"       // Importing the os package to interact with the operating system
	"strings"  // Importing the strings package for string manipulation
)

func main() {

	// Creating a scanner that reads from standard input
	scanner := bufio.NewScanner(os.Stdin)
	// Printing the CSV header
	fmt.Printf("domain, hasMX, hasSPF, sprRecord,hasDMARC,dmarcRecord\n")

	// Continuously scanning lines from standard input until an error or EOF
	for scanner.Scan() {
		// Getting the scanned text (domain) from the input
		domain := scanner.Text()
		// Calling the checkDomain function with the scanned domain as an argument
		checkDomain(domain)
	}

	// Handling errors from the scanner (if any)
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error could not read from input: %v", err)  // Using Fatalf to log error and exit
	}

}

// checkDomain function that checks for MX, SPF, and DMARC records for a given domain
func checkDomain(domain string) {

	// Initializing boolean variables to check the presence of MX, SPF, and DMARC records
	var hasMX, hasSPF, hasDMARC bool
	// Initializing strings to store the SPF and DMARC records
	var spfRecord, dmarcRecord string

	// Looking up MX records for the given domain
	mxRecords, err := net.LookupMX(domain)
	// Error handling for the MX lookup
	if err != nil {
		log.Printf("error: %v", err)
	}
	// If there are MX records, set hasMX to true
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// Looking up TXT records for the given domain
	txtRecords, err := net.LookupTXT(domain)
	// Error handling for the TXT lookup
	if err != nil {
		log.Printf("error: %v", err)
	}

	// Iterating over the TXT records to check for SPF and DMARC records
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record  // Storing the SPF record
		}
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record  // Storing the DMARC record
		}
	}

	// Looking up DMARC TXT records for the given domain with "_dmarc" prefix
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	// Error handling for the DMARC TXT lookup
	if err != nil {
		log.Printf("error: %v", err)
	}

	// Iterating over the DMARC TXT records to check and store DMARC records
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record  // Storing the DMARC record if it exists
			break  // Exiting the loop once the DMARC record is found
		}
	}

	// Printing the result as a CSV line
	fmt.Printf("%s,%t,%t,%s,%t,%s\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}

