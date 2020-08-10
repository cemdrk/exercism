// Package hamming provides a Hamming Distance function.
package hamming

import "errors"

// Distance returns distance between two strings of equal length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("string length must be equal")
	}

	diff := 0

	// count letter differences
	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}

	return diff, nil
}
