// Package twofer implements ShareWith function.
package twofer

import "fmt"

// ShareWith returns "One for (name), one for me.".
func ShareWith(name string) string {
	if (name == "") {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
