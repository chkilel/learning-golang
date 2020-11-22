// Package twofer iplement Two-fer.
// One for Exercism, one foe me.
package twofer

// ShareWith return Two-fer
func ShareWith(name string) string {

	if len(name) == 0 {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
