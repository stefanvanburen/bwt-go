package bwt

import (
	"fmt"
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

// NaiveIBWT returns the inverse of the Burrows-Wheeler Transform on a string
func NaiveIBWT(s string) string {
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

func IBWT(r string, il []int) string {
	s := ""
	i := il[0]
	fmt.Println(r)
	for z := 0; z < len(r); z++ {
		fmt.Println(i)
		s = s + string(r[i])
		i = il[i]
	}
	return s
}

// BWT creates an indexlist for use with IBWT, to lower the complexity of IBWT
func BWT(s string) (string, []int) {
	il := make([]int, len(s))
	x := getRotations(s)
	y := make([]string, len(x))
	copy(y, x)
	sort.Strings(y)
	fmt.Println(x)
	fmt.Println(y)
	l := len(s)
	for i, v := range y {
		i1 := index(v, x)
		//fmt.Println("i1 first: ", i1)
		if i1 < l-1 {
			i1++
		} else {
			i1 = 0
		}
		//fmt.Println("i1: ", i1)
		i2 := index(x[i1], y)
		//fmt.Println("i2: ", i2)
		il[i] = i2
		//fmt.Println(il)
	}
	fmt.Println(il)
	return takeLast(y), il
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
		fmt.Println(a)
		a = a[:j] + a[j+1:]
		fmt.Println(a)
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
	//fmt.Println(s)
	//fmt.Println(ss)
	for i, e := range ss {
		if strings.Compare(s, e) == 0 {
			//fmt.Println(i)
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
