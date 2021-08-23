package main

import (
	"flag"
	"fmt"
	"github.com/achenet/candy-store/calculator"
	"github.com/achenet/candy-store/client"
	"os"
	"sort"
)

const defaultURL = "http://candystore.zimpler.net"

func main() {

	alt := flag.Bool("alt", true, "Use alternate procedure to produce output")
	flag.Parse()

	page, err := client.GetHTMLPage(defaultURL)
	if err != nil {
		fmt.Println("Error getting web page:", err.Error())
		os.Exit(1)
	}

	if *alt {
		detailsTable := client.ParseHTMLPageForDetailsTable(page)
		summaryTable := calculator.CalculateTopCustomerFavorites(detailsTable)
		prettyPrintTable(summaryTable)
		return
	}

	summaryTable := client.ParseHTMLPageForSummaryTable(page)

	sort.Slice(summaryTable, func(i, j int) bool {
		return summaryTable[i].TotalSnacks > summaryTable[j].TotalSnacks
	})

	prettyPrintTable(summaryTable)
}

func prettyPrintTable(t []calculator.TopCustomerFavourite) {
	fmt.Println("[")
	for _, f := range t {
		fmt.Println("\t{")
		fmt.Printf("\t\tname: \"%s\",\n", f.Name)
		fmt.Printf("\t\tfavouriteSnack: \"%s\",\n", f.FavouriteSnack)
		fmt.Printf("\t\ttotalSnacks: %v,\n", f.TotalSnacks)
		fmt.Println("\t},")
	}

	fmt.Println("]")
}
