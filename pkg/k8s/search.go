package k8s

import (
	"regexp"
	"strings"
)

func normalizeName(name string) string {
	name = strings.ToLower(name)
	name = strings.TrimSpace(name)

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return name
	}

	return reg.ReplaceAllString(name, "")
}
