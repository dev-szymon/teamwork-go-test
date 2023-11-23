package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dev-szymon/teamwork-go-test/customerimporter"
)

func main() {

	f, err := os.Open("./customerimporter/customers.csv")
	if err != nil {
		log.Fatalf("Could not open file: %+v", err)
	}
	domains, err := customerimporter.GetDomainCounts(f)
	if err != nil {
		log.Fatalf("Error geting domains: %+v", err)
	}
	defer f.Close()

	fmt.Printf("%+v", domains)
}
