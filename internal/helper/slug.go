package helper

import (
	"regexp"
	"strings"
)

func GenerateSlug(s string) string {
	s = strings.ToLower(s)
	// Replace non-alphanumeric with hyphen
	reg := regexp.MustCompile("[^a-z0-9]+")
	s = reg.ReplaceAllString(s, "-")
	// Trim hyphens from both ends
	return strings.Trim(s, "-")
}
