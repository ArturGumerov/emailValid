package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Start system, write email..\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input: &v\n", err)
	}

}

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(email)
}

func checkDomain(domain string) {

	var hasMX, hasSPF, hasDMARC, isEmailval bool
	var spfRecord, dmarcRecord string

	isEmailval = isEmailValid(domain)
	domain = strings.Split(domain, "@")[1]

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, v := range txtRecords {
		if strings.HasPrefix(v, "v=spf1") {
			hasSPF = true
			spfRecord = v
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, v := range dmarcRecords {
		if strings.HasPrefix(v, "v=DMARC1") {
			dmarcRecord = v
			break
		}
	}

	fmt.Printf("isEmailvalid = %v, domain= %v , hasMX = %v, hasSPF= %v, spfRecord=%v, hasDMARC= %v, dmarcRecord=%v\n", isEmailval, domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

}
