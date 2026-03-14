package state

import "fmt"

import "web/crawl/internal/crawler/policy"

func InitState(seedURL string) (map[string]struct{}, []string, policy.RobotsRules) {
	seen := map[string]struct{}{seedURL: {}}
	queue := []string{seedURL}
	robotsRules := policy.ReadRobotsFile(seedURL)
	fmt.Println("Seed URL queued:", queue)

	return seen, queue, robotsRules
}

func Dequeue(queue []string) (string, []string) {
	current := queue[0]
	remaining := queue[1:]
	fmt.Println("Testing link:", current)
	return current, remaining
}

func EnqueueIfNew(queue []string, seen map[string]struct{}, normalized string) []string {
	if _, exists := seen[normalized]; exists {
		return queue
	}

	seen[normalized] = struct{}{}
	return append(queue, normalized)
}
