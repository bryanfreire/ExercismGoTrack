package scrabble

import "strings"

// Score the scrabble word
func Score(s string) (r int) {

	s = strings.ToUpper(s)
	for _, c := range s {
		switch c {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			r++
		case 'D', 'G':
			r += 2
		case 'B', 'C', 'M', 'P':
			r += 3
		case 'F', 'H', 'V', 'W', 'Y':
			r += 4
		case 'K':
			r += 5
		case 'J', 'X':
			r += 8
		case 'Q', 'Z':
			r += 10
		}
	}

	return
}
