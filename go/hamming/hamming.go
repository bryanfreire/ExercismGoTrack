package hamming

import "errors"

// Distance between dna strings
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("the strings must have the same length")
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
