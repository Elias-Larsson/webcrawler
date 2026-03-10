package crawler

import (
	"bufio"
	"fmt"
	"os"
)

func readRobotsFile() {
	file, err := os.Open("robots.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)

	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			fmt.Printf("ReadLine: %q\n", line)
		}
		if err != nil {
			panic(err)
		}
	}
}
 