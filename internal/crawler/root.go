package crawler

import (
	"time"
)

func RootExec(seedURL string) {
	started := time.Now()
	seen, queue := initState(seedURL)

	for len(queue) > 0 {
		var current string
		current, queue = dequeue(queue)

		htmlPage := extractHtmlPage(current)
		links := extractLinks(htmlPage)

		for _, link := range links {
			normalized, ok := normalizeLink(current, link)
			if !ok {
				continue
			}

			queue = enqueueIfNew(queue, seen, normalized)
		}
	}
	printSummary(seen, started)
}
