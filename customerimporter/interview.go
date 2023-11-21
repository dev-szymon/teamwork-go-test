// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

const emailColLabel = "email"

type Domain struct {
	Name  string
	Count int
}

func getDomainFromRow(row []string, emailColIndex int) string {
	if len(row) < emailColIndex || emailColIndex < 0 {
		return ""
	}

	emailParts := strings.Split(row[emailColIndex], "@")
	if len(emailParts) < 2 {
		return ""
	}
	return strings.ToLower(emailParts[1])
}

func GetDomainCounts(filepath string) ([]*Domain, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(bufio.NewReader(f))

	emailColIndex := -1
	headerRow, err := csvReader.Read()
	if err != nil {
		return nil, err
	}
	for i, label := range headerRow {
		if label == emailColLabel {
			emailColIndex = i
		}
	}
	if emailColIndex < 0 {
		return nil, fmt.Errorf("email column not found in file: %s", filepath)
	}

	domainCounts := map[string]int{}
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		domain := getDomainFromRow(row, emailColIndex)
		if domain == "" {
			fmt.Printf("Could not extract domain name from row: %+v\n", row)
			continue
		}
		domainCounts[domain]++
	}

	domains := []*Domain{}
	for name, count := range domainCounts {
		domains = append(domains, &Domain{Name: name, Count: count})
	}
	sort.Slice(domains, func(a, b int) bool {
		return domains[a].Name < domains[b].Name
	})

	return domains, nil
}
