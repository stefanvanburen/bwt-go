package BWT

import (
	"sort"
	"strings"
)

const eofChar = "$"

func rotate(s string) string {
	return string(s[len(s)-1]) + string(s[:len(s)-1])
}

func getRotations(s string) (ss []string) {
	for i := 0; i < len(s); i++ {
		ss = append(ss, s)
		s = rotate(s)
	}
	return ss
}

func takeLast(ss []string) string {
	z := ""
	for _, s := range ss {
		z += string(s[len(s)-1])
	}
	return z
}

func index(s string, ss []string) int {
	for i, e := range ss {
		if strings.Compare(s, e) == 0 {
			return i
		}
	}
	return -1
}

func findLast(ss []string) string {
	for _, e := range ss {
		if string(e[len(e)-1]) == "$" {
			// Return the string up until the eofChar
			return string(e[:len(e)-1])
		}
	}
	return ""
}

// BWT returns the result of the Burrows-Wheeler Transform on a string
func BWT(s string) string {
	s = s + eofChar
	x := getRotations(s)
	sort.Strings(x)
	return takeLast(x)
}

// IBWT returns the inverse of the Burrows-Wheeler Transform on a string
func IBWT(s string) string {
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
