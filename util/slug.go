package util

import (
	"regexp"
	"strings"
)

func Slugify(input string) string {
	slug := strings.ToLower(input)

	slug = strings.ReplaceAll(slug, " ", "-")

	reg := regexp.MustCompile(`[^a-z0-9-]`)
	slug = reg.ReplaceAllString(slug, "")

	slug = strings.Trim(slug, "-")

	return slug
}
