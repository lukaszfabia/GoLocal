package normalizer

import (
	"regexp"
	"strings"
)

func clean(e string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9_\s]+`)
	//remove special chars
	cleaned := re.ReplaceAllString(e, " ")

	// remove many spaces
	cleaned = regexp.MustCompile(`\s+`).ReplaceAllString(cleaned, " ")

	// remove trailing space
	cleaned = strings.TrimSpace(cleaned)

	// set floor as a default
	return strings.ReplaceAll(cleaned, " ", "_")
}

/*
Normalizes list of strings

Example:
  - ['  tt#sa12', 'tag%^test'] -> ['tt_sa12', 'tag_test']

Params:

  - strings to normalize

Returns:

  - normalized list
*/
func Normalizer(lst []string) []string {
	res := []string{}

	for _, e := range lst {
		res = append(res, clean(e))
	}

	return res
}
