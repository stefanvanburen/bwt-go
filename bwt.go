// Package bwt implements variants of the [Burrows-Wheeler transform].
//
// [Burrows-Wheeler transform]: https://en.wikipedia.org/wiki/Burrows%E2%80%93Wheeler_transform
package bwt

import (
	"slices"
	"strings"
)

// EOFChar is used by the BWT and InverseBWT functions to mark the end of the
// input string. It must be unique within the string.
var EOFChar = "$"

// BWT returns the result of the Burrows-Wheeler Transform on the input string,
// using the EOFChar as a marker.
func BWT(input string) string {
	input = input + EOFChar
	rotations := getRotations(input)
	slices.Sort(rotations)
	return takeLast(rotations)
}

// InverseBWT returns the inverse of the Burrows-Wheeler Transform on the input
// string, using the EOFChar as a marker.
func InverseBWT(input string) string {
	characters := strings.Split(input, "")

	z := make([]string, len(input))

	for i := 0; i < len(characters); i++ {
		for j := 0; j < len(characters); j++ {
			z[j] = characters[j] + z[j]
		}
		slices.Sort(z)
	}

	return findLast(z)
}

// rotate returns a rotated version of s where the last character becomes
// the first.
func rotate(s string) string {
	return string(s[len(s)-1]) + string(s[:len(s)-1])
}

// getRotations returns all rotations of s.
func getRotations(s string) []string {
	rotations := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		rotations[i] = s
		s = rotate(s)
	}
	return rotations
}

// takeLast returns a string containing the last character of all of the
// strings in ss.
func takeLast(ss []string) (lastCharacters string) {
	for _, s := range ss {
		lastCharacters += string(s[len(s)-1])
	}
	return lastCharacters
}

// findLast returns the string in ss that has EOFChar as a suffix.
func findLast(ss []string) string {
	for _, str := range ss {
		if strings.HasSuffix(str, EOFChar) {
			return strings.TrimSuffix(str, EOFChar)
		}
	}
	return ""
}
