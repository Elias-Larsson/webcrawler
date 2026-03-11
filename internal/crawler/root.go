package crawler

import (
	"time"
)

type crawlResult struct {
	url   string
	links []string
}

func RootExec(seedURL string) {
	started := time.Now()
	const numWorkers = 10

	jobs := make(chan string)
	results := make(chan *crawlResult, 8)

	seen, queue, robotsRules := initState(seedURL)

	for i := 0; i < numWorkers; i++ {
		go worker(jobs, results)
	}

	pending := 0

	var seed string
	seed, queue = dequeue(queue)
	jobs <- seed
	pending++

	for pending > 0 {
		result := <-results
		pending--

		for _, link := range result.links {
			normalized, ok := normalizeLink(result.url, link)
			if !ok {
				continue
			}
			if !isAllowed(robotsRules, normalized) {
				continue
			}
			queue = enqueueIfNew(queue, seen, normalized)
		}

		for len(queue) > 0 && pending < numWorkers {
			var next string
			next, queue = dequeue(queue)
			jobs <- next
			pending++
		}
	}

	close(jobs)
	printSummary(seen, started)
}

func worker(jobs <-chan string, results chan<- *crawlResult) {
	for url := range jobs {
		page := extractHtmlPage(url)
		links := extractLinks(page)
		results <- &crawlResult{url: url, links: links}
	}
}
