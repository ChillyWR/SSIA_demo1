package pkg

import "strings"

// CompareStrings check if strings are equal
// String trimming and cleanup
func CompareStrings(input, answer string) bool {
	input = strings.Join(strings.Fields(input), " ")
	return strings.EqualFold(input, answer)
}
