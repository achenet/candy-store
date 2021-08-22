package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const defaultURL = "http://candystore.zimpler.net"

func main() {
	page, err := client.GetHMTLPage(defaultURL)
	if err != nil {
		fmt.Println("Error getting web page:", err.Error())
		os.Exit(1)
	}

	summaryTable := ParseHTMLPageForSummaryTable(page)

	// Marshal to JSON and print to get requested format
	bytes, err := json.Marhsal(summaryTable)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err.Error())
		os.Exit(1)
	}

	fmt.Println(bytes)
}
