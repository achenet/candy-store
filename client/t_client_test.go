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
    fmt.Println("Test summary table")
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
		t.Error("Got:", got, "\nExpected:", expectedSummaryTable)
	}
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


func TestParseHTMLPageForDetailsTable(t *testing.T) {
    fmt.Println("Test details table")
	page, err := GetHTMLPage("https://candystore.zimpler.net/")
	if err != nil {
		t.Error(err.Error())
	}
    
    expectedDetailsTable := []calculator.CustomerEntry{
        {
            Name: "Annika",
            Candy: "Geisha",
            Eaten: 100,
        },
        {
            Name: "Jonas",
            Candy: "Geisha",
            Eaten: 200,
        },
        {
            Name: "Jonas",
            Candy: "Kexchoklad",
            Eaten: 100,
        },
        {
            Name: "Aadya",
            Candy: "Nötchoklad",
            Eaten: 2,
        },
        {
            Name: "Jonas",
            Candy: "Nötchoklad",
            Eaten: 3,
        },
        {
            Name: "Jane",
            Candy: "Nötchoklad",
            Eaten: 17,
        },
        {
            Name: "Annika",
            Candy: "Geisha",
            Eaten: 100,
        },
        {
            Name: "Jonas",
            Candy: "Geisha",
            Eaten: 700,
        },
        {
            Name: "Jane",
            Candy: "Nötchoklad",
            Eaten: 4,
        },
        {
            Name: "Aadya",
            Candy: "Center",
            Eaten: 7,
        },
        {
            Name: "Jonas",
            Candy: "Geisha",
            Eaten: 900,
        },
        {
            Name: "Jane",
            Candy: "Nötchoklad",
            Eaten: 1,
        },
        {
            Name: "Jonas",
            Candy: "Kexchoklad",
            Eaten: 12,
        },
        {
            Name: "Jonas",
            Candy: "Plopp",
            Eaten: 40,
        },
        {
            Name: "Jonas",
            Candy: "Center",
            Eaten: 27,
        },
        {
            Name: "Aadya",
            Candy: "Center",
            Eaten: 2,
        },
        {
            Name: "Annika",
            Candy: "Center",
            Eaten: 8,
        },
    }

    got := ParseHTMLPageForDetailsTable(page)
	if !areEqualCustomerEntry(expectedDetailsTable, got) {
		t.Error("Got:", got, "\nExpected:", expectedDetailsTable)
	}
    fmt.Println(got)
}

func areEqualCustomerEntry(first, second []calculator.CustomerEntry) bool {
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
