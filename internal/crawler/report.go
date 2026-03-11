package crawler

import (
	"fmt"
	"time"
)

func printSummary(seen map[string]struct{}, started time.Time) {
	duration := time.Since(started)

	fmt.Println()
	fmt.Printf("%v seen links under the duration %v", len(seen), duration)
	fmt.Println()
	fmt.Println("Seen links:")
	for link := range seen {
		fmt.Printf("- %s\n", link)
	}
}
