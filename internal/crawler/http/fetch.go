package crawlhttp

import (
	"net/http"

	"golang.org/x/net/html"
)

func ExtractHTMLPage(url string) *html.Node {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		panic(err)
	}

	return doc
}
