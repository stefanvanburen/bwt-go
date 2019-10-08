package bwt

import (
	"sort"
	"strings"
)

const eofChar = "$"

// NaiveBWT returns the result of the Burrows-Wheeler Transform on a string
func NaiveBWT(s string) string {
	s = s + eofChar
	x := getRotations(s)
	sort.Strings(x)
	return takeLast(x)
}

// NaiveInverseBWT returns the inverse of the Burrows-Wheeler Transform on a string
func NaiveInverseBWT(s string) string {
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

// IndexList is used to lower the complexity of the InverseBWT operation.
type IndexList []int

// BWT creates an indexlist for use with IBWT, to lower the complexity of IBWT
func BWT(s string) (string, IndexList) {
	x := getRotations(s)
	y := make([]string, len(x))
	copy(y, x)
	sort.Strings(y)
	il := make(IndexList, len(s))
	l := len(s)
	for i, v := range y {
		i1 := index(v, x)
		if i1 < l-1 {
			i1++
		} else {
			i1 = 0
		}
		i2 := index(x[i1], y)
		il[i] = i2
	}
	return takeLast(y), il
}

// InverseBWT returns the inverse of a BWT transformation, using an index list.
func InverseBWT(r string, il IndexList) string {
	s := ""
	i := il[0]
	for z := 0; z < len(r); z++ {
		s = s + string(r[i])
		i = il[i]
	}
	return s
}

// MTF returns the Move-To-Front Coding of a string
// Typically used on a Burrows-Wheeler Transformed string
func MTF(s string) []int {
	a := alphabet(s)
	mtf := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		// Find the index of s[i]
		j := strings.Index(a, string(s[i]))
		mtf = append(mtf, j)
		// Move it to the front
		char := string(a[j])
		a = a[:j] + a[j+1:]
		a = char + a
	}
	return mtf
}

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
		if string(e[len(e)-1]) == eofChar {
			// Return the string up until the eofChar
			return string(e[:len(e)-1])
		}
	}
	return ""
}

// Get a sorted alphabet of all the characters in the string
func alphabet(s string) string {
	set := make(map[string]bool)
	for i := 0; i < len(s); i++ {
		set[string(s[i])] = true
	}
	ss := make([]string, len(s))
	for k := range set {
		ss = append(ss, k)
	}
	sort.Strings(ss)
	r := ""
	for _, v := range ss {
		r += v
	}
	return r
}
