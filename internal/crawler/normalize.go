package crawler

import "net/url"

func normalizeLink(currentPage string, rawLink string) (string, bool) {
	base, err := url.Parse(currentPage)
	if err != nil {
		return "", false
	}

	u, err := url.Parse(rawLink)
	if err != nil {
		return "", false
	}

	resolved := base.ResolveReference(u)
	if resolved.Scheme != "http" && resolved.Scheme != "https" {
		return "", false
	}

	resolved.Fragment = ""
	return resolved.String(), true
}
