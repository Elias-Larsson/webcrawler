package crawler

import (
	"fmt"
	"net/url"
	"time"
)

func Exec() {
	seedURL := "http://localhost:5500"
	timer := time.Now()

	seen := make(map[string]struct{})
	queue := []string{seedURL}
	seen[seedURL] = struct{}{}

	fmt.Println("Seed URL queued:", queue)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Println("Testing link:", current)

		base, err := url.Parse(current)
		if err != nil {
			continue
		}

		htmlPage := extractHtmlPage(current)
		links := extractLinks(htmlPage)

		for _, link := range links {
			u, err := url.Parse(link)
			if err != nil {
				continue
			}

			resolved := base.ResolveReference(u)

			if resolved.Scheme != "http" && resolved.Scheme != "https" {
				continue
			}

			resolved.Fragment = ""

			normalized := resolved.String()
			if _, exists := seen[normalized]; exists {
				continue
			}
			seen[normalized] = struct{}{}
			queue = append(queue, normalized)
		}
	}
	duration := time.Since(timer)
	fmt.Println()
	fmt.Printf("%v seen links under the duration %v", len(seen), duration)
	fmt.Println()
	fmt.Println("Seen links:")
for link := range seen {
    fmt.Printf("- %s\n", link)
}

}
