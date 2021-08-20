package client

import (
	"github.com/achenet/candy-store/calculator"
	"io"
	"net/http"

	"golang.org/x/net/html"
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

func ParseHTMLPageAndCreateTable(webpage string) []calculator.CustomerEntry {

    r := strings.NewReader(webpage)
    z := html.NewTokenizer(r)
    _ = z
	return []calculator.CustomerEntry{}
}
