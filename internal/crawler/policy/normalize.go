package policy

import "net/url"

func NormalizeLink(currentPage string, rawLink string, seedHostname string) (string, bool) {
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

	if resolved.Hostname() != seedHostname {
		return "", false
	}

	return resolved.String(), true
}
