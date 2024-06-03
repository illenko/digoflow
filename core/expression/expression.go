package expression

import (
	"regexp"
	"strings"
)

// placeholderPattern is the regular expression pattern used to match placeholders in an expression.
// Placeholders are expected to be in the format {{ placeholder }}.
const placeholderPattern = `{{\s*[a-zA-Z0-9\.]*\s*}}`

// cutset is the set of characters to remove from the placeholders.
const cutset = "{} "

var placeholderRegex = regexp.MustCompile(placeholderPattern)

// GetPlaceholders returns a list of placeholders found in the given expression.
// Placeholders are expected to be in the format {{ placeholder }}.
// The function will return a list of placeholders without the curly braces.
func GetPlaceholders(expression string) []string {
	placeholders := make([]string, 0)

	for _, match := range placeholderRegex.FindAllString(expression, -1) {
		placeholders = append(placeholders, strings.Trim(match, cutset))
	}

	return placeholders
}

// ReplacePlaceholders replaces placeholders in the given expression with the values provided.
// Placeholders are expected to be in the format {{ placeholder }}.
// The function will return the expression with the placeholders replaced by the values.
func ReplacePlaceholders(expression string, values map[string]string) string {
	return placeholderRegex.ReplaceAllStringFunc(expression, func(match string) string {
		placeholder := strings.Trim(match, cutset)
		return values[placeholder]
	})
}
