package crawler

import (
	crawlhttp "web/crawl/internal/crawler/http"
	"web/crawl/internal/crawler/parse"
)

func worker(jobs <-chan string, results chan<- *crawlResult) {
	for url := range jobs {
		page := crawlhttp.ExtractHTMLPage(url)
		links := parse.ExtractLinks(page)
		results <- &crawlResult{url: url, links: links}
	}
}