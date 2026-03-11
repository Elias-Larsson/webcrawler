package crawler

import (
	"bufio"
	"net/http"
	"net/url"
	"strings"
	"web/crawl/internal/utils"
)

type RobotsRules struct {
	Disallowed []string
	Allowed    []string
}

func ReadRobotsFile(baseURL string) RobotsRules {
	rules := RobotsRules{}
	res, err := http.Get(baseURL + "/robots.txt")

	if err != nil || res.StatusCode != 200 {
		return rules
	}

	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "Disallow:") {
			path := utils.TrimPrefix(line, "Disallow:")

			if path != "" {
				rules.Disallowed = append(rules.Disallowed, path)
			}

		} else if strings.HasPrefix(line, "Allow:") {
			path := utils.TrimPrefix(line, "Allow:")

			if path != "" {
				rules.Allowed = append(rules.Allowed, path)
			}

		}

	}
	
	return rules
}

func isAllowed(rules RobotsRules, rawURL string) bool {
	u, err := url.Parse(rawURL)

	if err != nil {
		return false
	}

	path := u.Path
	if path == "" {
		path = "/"
	}

	bestLen := -1
	bestAllow := true

	for _, a := range rules.Allowed {
		if strings.HasPrefix(path, a) && len(a) > bestLen {
			bestLen = len(a)
			bestAllow = true
		}
	}

	for _, d := range rules.Disallowed {
		if strings.HasPrefix(path, d) && len(d) > bestLen {
			bestLen = len(d)
			bestAllow = false
		}
	}

	return bestAllow
}
