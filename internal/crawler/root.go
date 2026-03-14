package crawler

import (
	"net/url"
	"time"

	"web/crawl/internal/crawler/policy"
	"web/crawl/internal/crawler/report"
	"web/crawl/internal/crawler/state"
)

type crawlResult struct {
	url   string
	links []string
}

func RootExec(seedURL string) {
	started := time.Now()
	const numWorkers = 10

	seedParsed, _ := url.Parse(seedURL)
	seedHostname := seedParsed.Hostname()

	jobs := make(chan string)
	results := make(chan *crawlResult, 8)

	seen, queue, robotsRules := state.InitState(seedURL)

	for range numWorkers {
		go worker(jobs, results)
	}

	pending := 0

	var seed string
	seed, queue = state.Dequeue(queue)
	jobs <- seed
	pending++

	for pending > 0 {
		result := <-results
		pending--
		for _, link := range result.links {
			normalized, ok := policy.NormalizeLink(result.url, link, seedHostname)
			if !ok {
				continue
			}
			if !policy.IsAllowed(robotsRules, normalized) {
				continue
			}
			queue = state.EnqueueIfNew(queue, seen, normalized)
		}

		for len(queue) > 0 && pending < numWorkers {
			var next string
			next, queue = state.Dequeue(queue)
			jobs <- next
			pending++
		}
	}

	close(jobs)
	report.PrintSummary(seen, started)
}


