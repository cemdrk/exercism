// Package raindrops provides raindrop functions
package raindrops

import "strconv"

var rules = []struct {
	factor int
	sound  string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

// Convert converts a number to raindrop sound
func Convert(drop int) string {
	result := ""

	for _, rule := range rules {
		if drop%rule.factor == 0 {
			result += rule.sound
		}
	}

	if len(result) == 0 {
		result += strconv.Itoa(drop)
	}

	return result
}
