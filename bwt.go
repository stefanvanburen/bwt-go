// Package bwt implements variants of the Burrows-Wheeler transform algorithm [1].
//
// [1]: https://en.wikipedia.org/wiki/Burrows%E2%80%93Wheeler_transform
package bwt

import (
	"sort"
	"strings"
)

var EOFChar = "$"

// BWT returns the result of the Burrows-Wheeler Transform on a string, using
// the EOFChar as a marker.
func BWT(s string) string {
	s = s + EOFChar
	x := getRotations(s)
	sort.Strings(x)
	return takeLast(x)
}

// InverseBWT returns the inverse of the Burrows-Wheeler Transform on a string,
// using the EOFChar as a marker.
func InverseBWT(s string) string {
	x := strings.Split(s, "")
	z := make([]string, len(s))
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x); j++ {
			z[j] = x[j] + z[j]
		}
		sort.Strings(z)
	}
	return findLast(z)
}

func rotate(s string) string {
	return string(s[len(s)-1]) + string(s[:len(s)-1])
}

func getRotations(s string) []string {
	rotations := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		rotations[i] = s
		s = rotate(s)
	}
	return rotations
}

func takeLast(ss []string) string {
	z := ""
	for _, s := range ss {
		z += string(s[len(s)-1])
	}
	return z
}

func findLast(ss []string) string {
	for _, e := range ss {
		if string(e[len(e)-1]) == EOFChar {
			// Return the string up until the eofChar
			return string(e[:len(e)-1])
		}
	}
	return ""
}
