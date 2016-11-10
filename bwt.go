package BWT

import (
	"sort"
	"strings"
)

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

// BWT returns the the Burrows-Wheeler Transform on a string
func BWT(s string) (string, []int) {
	x := getRotations(s)
	y := x
	sort.Strings(y)
	var idxlist []int
	for _, e := range y {
		i1 := index(e, x)
		if i1 < len(s)-1 {
			i1++
		} else {
			i1 = 0
		}
		i2 := index(x[i1], y)
		idxlist = append(idxlist, i2)
	}
	return takeLast(x), idxlist
}

// IBWT returns the inverse of the Burrows-Wheeler Transform on a string
func IBWT(s string, il []int) string {
	t := ""
	x := il[0]
	for i := 0; i < len(s); i++ {
		t += string(s[x])
		x = il[x]
	}
	return t
}
