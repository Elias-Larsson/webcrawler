package crawler

import (
	"bufio"
	"net/http"
)

func ReadRobotsFile(base string) []string {
	resultFile := []string{}
	file, err := http.Get(base + "/robots.txt")
	if err != nil {
		panic(err)
	}
	defer file.Body.Close()

	scanner := bufio.NewScanner(file.Body)
	for scanner.Scan() {
		resultFile = append(resultFile, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return resultFile
}
