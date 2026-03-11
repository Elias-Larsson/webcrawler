package crawler

import (
	"net/http"

	"golang.org/x/net/html"
)

func extractHtmlPage(url string) *html.Node {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	return doc
}
