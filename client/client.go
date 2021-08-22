package client

import (
	"fmt"
	"github.com/achenet/candy-store/calculator"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func GetHTMLPage(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func ParseHTMLPageForSummaryTable(webpage string) []calculator.TopCustomerFavourite {
	summaryTable := make([]calculator.TopCustomerFavourite, 0)
	r := strings.NewReader(webpage)

	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "table" {
			if contains(n.Attr, html.Attribute{
				Key: "class",
				Val: "top.customers summary",
			}) {
				summaryTable = extractSummaryTable(n)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)
	return summaryTable
}

func contains(attrList []html.Attribute, wantedAttr html.Attribute) bool {
	for _, a := range attrList {
		if a == wantedAttr {
			return true
		}
	}
	return false
}

func extractSummaryTable(n *html.Node) []calculator.TopCustomerFavourite {
	summaryList := make([]calculator.TopCustomerFavourite, 0)
	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.TextNode && n.Parent.Data == "td" {
			if hasAttributeKey(n.Parent.Attr, "x-total-candy") {
				summaryList = append(summaryList, calculator.TopCustomerFavourite{
					Name:        n.Data,
					TotalSnacks: getTotalSnacksFromAttr(n.Parent.Attr),
				})
			} else {
				summaryList[len(summaryList)-1].FavouriteSnack = n.Data
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)
	return summaryList
}

func getTotalSnacksFromAttr(attr []html.Attribute) int {
	for _, a := range attr {
		if a.Key == "x-total-candy" {
			num, err := strconv.Atoi(a.Val)
			if err != nil {
				return -1
			}
			return num
		}
	}
	return 0
}

func hasAttributeKey(attrList []html.Attribute, key string) bool {
	for _, a := range attrList {
		if a.Key == key {
			return true
		}
	}
	return false
}

func ParseHTMLPageForDetailsTable(webpage string) []calculator.CustomerEntry {
	detailsTable := make([]calculator.CustomerEntry, 0)
	r := strings.NewReader(webpage)

	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "table" {
			if contains(n.Attr, html.Attribute{
				Key: "class",
				Val: "top.customers details",
			}) {
				detailsTable = extractDetailsTable(n)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)
	return detailsTable
}

func extractDetailsTable(n *html.Node) []calculator.CustomerEntry {
	detailsTable := make([]calculator.CustomerEntry, 0)

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "tr" {
			detailsTable = append(detailsTable, createCustomerEntry(n))
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)
	return detailsTable
}

func createCustomerEntry(n *html.Node) calculator.CustomerEntry {
	// We know that n is a row
	// As such, it has 3 <td> children, each of which has one text child.
	i := 0
	entry := calculator.CustomerEntry{}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode && n.Parent.Data == "td" && n.Parent.Type == html.ElementNode && n.Data != "" {
			switch i {
			case 0:
				entry.Name = n.Data
				i++
			case 1:
				entry.Candy = n.Data
				i++
			case 2:
				num, _ := strconv.Atoi(n.Data)
				entry.Eaten = num
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)
	fmt.Println("adding entry:", entry)
	return entry
}
