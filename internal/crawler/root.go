package crawler

import (
	"net/url"
	"time"
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

	seen, queue, robotsRules := initState(seedURL)

	for range numWorkers {
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
			normalized, ok := normalizeLink(result.url, link, seedHostname)
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


