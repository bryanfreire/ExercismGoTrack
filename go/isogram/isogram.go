package isogram

import "strings"

// IsIsogram strings
func IsIsogram(s string) (r bool) {

	s = strings.ToLower(s)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > 64 && s[i] < 123 && s[i] == s[j] {
				return false
			}
		}
	}
	return true
}
