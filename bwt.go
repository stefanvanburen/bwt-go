// Package bwt implements variants of the [Burrows-Wheeler transform].
//
// [Burrows-Wheeler transform]: https://en.wikipedia.org/wiki/Burrows%E2%80%93Wheeler_transform
package bwt

import (
	"fmt"
	"slices"
	"strings"
)

// BWT returns the result of the Burrows-Wheeler Transform on the input string,
// using the eofCharacter as a marker. The eofCharacter cannot exist in the
// input string.
func BWT(input string, eofCharacter rune) (string, error) {
	if strings.ContainsRune(input, eofCharacter) {
		return "", fmt.Errorf("input contains eofCharacter")
	}
	input += string(eofCharacter)
	rotations := getRotations(input)
	slices.Sort(rotations)
	var lastCharacters string
	for _, rotation := range rotations {
		lastCharacters += string(rotation[len(rotation)-1])
	}
	return lastCharacters, nil
}

// InverseBWT returns the inverse of the Burrows-Wheeler Transform on the input
// string, using the eofCharacter as a marker.
func InverseBWT(input string, eofCharacter rune) (string, error) {
	characters := strings.Split(input, "")

	z := make([]string, len(input))

	for range len(characters) {
		for j := range len(characters) {
			z[j] = characters[j] + z[j]
		}
		slices.Sort(z)
	}

	for _, str := range z {
		if strings.HasSuffix(str, string(eofCharacter)) {
			return strings.TrimSuffix(str, string(eofCharacter)), nil
		}
	}
	return "", fmt.Errorf("did not find shift with eofCharacter as a suffix - invariant broken")
}

// rotate returns a rotated version of s where the last character becomes
// the first.
func rotate(s string) string {
	return string(s[len(s)-1]) + string(s[:len(s)-1])
}

// getRotations returns all rotations of s.
func getRotations(s string) []string {
	rotations := make([]string, len(s))
	for i := range len(s) {
		rotations[i] = s
		s = rotate(s)
	}
	return rotations
}
