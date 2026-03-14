package utils

import "strings"

func TrimPrefix(line string, prefix string) string {
	path := strings.TrimSpace(strings.TrimPrefix(line, prefix))

	return path
}