package main

import (
	"fmt"
	"os"
    "github.com/achenet/candy-store/client"
    "github.com/achenet/candy-store/calculator"
    "sort"
)

const defaultURL = "http://candystore.zimpler.net"

func main() {
	page, err := client.GetHTMLPage(defaultURL)
	if err != nil {
		fmt.Println("Error getting web page:", err.Error())
		os.Exit(1)
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
