// Package bwt implements variants of the [Burrows-Wheeler transform].
//
// [Burrows-Wheeler transform]: https://en.wikipedia.org/wiki/Burrows%E2%80%93Wheeler_transform
package bwt

import (
	"fmt"
	"slices"
	"strings"
	"unicode/utf8"
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
	var buf strings.Builder
	buf.Grow(len(input))
	for _, rotation := range rotations {
		lastRune, _ := utf8.DecodeLastRuneInString(rotation)
		buf.WriteRune(lastRune)
	}
	return buf.String(), nil
}

// InverseBWT returns the inverse of the Burrows-Wheeler Transform on the input
// string, using the eofCharacter as a marker.
func InverseBWT(input string, eofCharacter rune) (string, error) {
	characters := strings.Split(input, "")
	z := make([]string, len(characters))

	for range len(characters) {
		for j := range len(characters) {
			z[j] = characters[j] + z[j]
		}
		slices.Sort(z)
	}

	for _, str := range z {
		if before, ok := strings.CutSuffix(str, string(eofCharacter)); ok {
			return before, nil
		}
	}
	return "", fmt.Errorf("did not find shift with eofCharacter as a suffix - invariant broken")
}

// rotate returns a rotated version of s where the last rune becomes
// the first.
func rotate(s []rune) []rune {
	rotated := make([]rune, len(s))
	rotated[0] = s[len(s)-1]
	copy(rotated[1:], s[:len(s)-1])
	return rotated
}

// getRotations returns all rotations of s.
func getRotations(s string) []string {
	runes := []rune(s)
	rotations := make([]string, len(runes))
	for i := range len(runes) {
		rotations[i] = string(runes)
		runes = rotate(runes)
	}
	return rotations
}
