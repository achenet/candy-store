package client

import (
	"fmt"
	"github.com/achenet/candy-store/calculator"
	"testing"
)

func TestGetHTMLPage(t *testing.T) {
	page, err := GetHTMLPage("https://candystore.zimpler.net/")
	if err != nil {
		t.Error(err.Error())
	}

	_ = page
	// fmt.Println(page)
}

func TestParseHTMLPageForSummaryTable(t *testing.T) {
	page, err := GetHTMLPage("https://candystore.zimpler.net/")
	if err != nil {
		t.Error(err.Error())
	}

	expectedSummaryTable := []calculator.TopCustomerFavourite{
		{
			Name:           "Aadya",
			FavouriteSnack: "Center",
			TotalSnacks:    11,
		},
		{
			Name:           "Annika",
			FavouriteSnack: "Geisha",
			TotalSnacks:    208,
		},
		{
			Name:           "Jonas",
			FavouriteSnack: "Geisha",
			TotalSnacks:    1982,
		},
		{
			Name:           "Jane",
			FavouriteSnack: "Nötchoklad",
			TotalSnacks:    22,
		},
	}

	got := ParseHTMLPageForSummaryTable(page)
	if !areEqual(expectedSummaryTable, got) {
		t.Error("Got:", got, "Expected", expectedSummaryTable)
	}
	fmt.Println(got)
}

func areEqual(first, second []calculator.TopCustomerFavourite) bool {
	if len(first) != len(second) {
		return false
	}
	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}
