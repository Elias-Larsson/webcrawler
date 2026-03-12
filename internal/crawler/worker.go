package crawler

func worker(jobs <-chan string, results chan<- *crawlResult) {
	for url := range jobs {
		page := extractHtmlPage(url)
		links := extractLinks(page)
		results <- &crawlResult{url: url, links: links}
	}
}