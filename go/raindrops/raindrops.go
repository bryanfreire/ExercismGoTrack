package raindrops

import "strconv"

// Convert int to string based on modulos
func Convert(x int) (r string) {

	if x%3 == 0 {
		r += "Pling"
	}
	if x%5 == 0 {
		r += "Plang"
	}
	if x%7 == 0 {
		r += "Plong"
	}
	if r == "" {
		r = strconv.Itoa(x)
	}

	return
}
