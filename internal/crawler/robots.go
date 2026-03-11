package crawler

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type RobotsRules struct {
	Disallowed []string
	Allowed    []string
}

func readRobotsFile(baseURL string) RobotsRules {
	applies := false
	crawlerUA := "aroez-agent"
	rules := RobotsRules{}
	res, err := http.Get(baseURL + "/robots.txt")

	if err != nil || res.StatusCode != 200 {
		return rules
	}

	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		key, value, ok := strings.Cut(line, ":")
		if !ok || value == "" || key == "" {
			continue
		}

		key, value = strings.ToLower(strings.TrimSpace(key)), strings.TrimSpace(value)

		switch key {
		case "user-agent":
			applies = value == "*" || strings.EqualFold(value, crawlerUA)

		case "disallow":
			if applies {
				rules.Disallowed = append(rules.Disallowed, value)
			}

		case "allow":
			if applies {
				rules.Allowed = append(rules.Allowed, value)
			}
		}
	}
	fmt.Println("robots ruleset: ", rules)
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
