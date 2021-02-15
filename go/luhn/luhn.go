package luhn

import (
	"strings"
)

// Valid lunh string
func Valid(input string) (valid bool) {
	input = strings.ReplaceAll(input, " ", "")
	if len(input) == 1 {
		return false
	}
	var total, current int
	var parity = len(input) % 2
	for i, n := range input {
		if n >= 48 && n <= 57 {
			current = int(n) - 48
			if i%2 == parity {
				current *= 2
				if current >= 10 {
					current -= 9
				}
			}
			total += current
		} else {
			total = -1
			break
		}
	}
	if total%10 == 0 {
		valid = true
	}
	return
}
