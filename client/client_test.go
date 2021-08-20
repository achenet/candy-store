package client

import (
	"fmt"
	"testing"
)

func TestGetHTMLPage(t *testing.T) {
	page, err := GetHTMLPage("https://candystore.zimpler.net/")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(page)
}

func TestParseHTMLPageAndCreateTable(t *testing.T) {
	page, err := GetHTMLPage("https://candystore.zimpler.net/")
	if err != nil {
		t.Error(err.Error())
	}

    fmt.Println(ParseHTMLPageAndCreateTable(page))
}
