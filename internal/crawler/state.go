package crawler

import "fmt"

func initState(seedURL string) (map[string]struct{}, []string, RobotsRules) {
	seen := map[string]struct{}{seedURL: {}}
	queue := []string{seedURL}
	robotsRules := readRobotsFile(seedURL)
	fmt.Println("Seed URL queued:", queue)

	return seen, queue, robotsRules
}

func dequeue(queue []string) (string, []string) {
	current := queue[0]
	remaining := queue[1:]
	fmt.Println("Testing link:", current)
	return current, remaining
}

func enqueueIfNew(queue []string, seen map[string]struct{}, normalized string) []string {
	if _, exists := seen[normalized]; exists {
		return queue
	}

	seen[normalized] = struct{}{}
	return append(queue, normalized)
}
