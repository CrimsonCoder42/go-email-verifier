# Domain Records Checker

This small Go application checks for the presence of MX, SPF, and DMARC records for a given domain and prints the results.

## How It Works

1. **User Input:** The program reads domains line-by-line from the standard input.
2. **Domain Check:** For each domain, it checks for the presence of MX, SPF, and DMARC records.
3. **Results Output:** It prints the result for each domain in CSV format.

## Code Overview

### main function (in the `main` package)

- Initializes a Scanner that reads from standard input.
- Prints CSV headers to the console.
- Continuously scans lines from standard input, and for each line, it calls the `checkDomain` function.
- If there is an error during scanning, it logs the error and exits.

### checkDomain function

- Initializes variables to check and store the presence and values of MX, SPF, and DMARC records.
- Uses the `net.LookupMX` function to look up MX records for a given domain.
- Uses the `net.LookupTXT` function to look up TXT records for a given domain.
- Iterates over TXT records to identify and store SPF and DMARC records.
- Uses `net.LookupTXT` again with a "\_dmarc" prefix to look up DMARC TXT records and stores them.
- Prints the result as a CSV line to the console.

## Usage

1. Run the Go application.
2. Enter the domains line-by-line in the console.
3. For each domain, the program will print a line with the results in CSV format.
