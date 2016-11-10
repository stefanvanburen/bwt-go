package BWT

import "testing"

var BWTtests = []struct {
	in  string
	out string
}{
	{"hey", "hye"},
	{"banana", "nnbaaa"},
}

func TestBWT(t *testing.T) {
	for _, tt := range BWTtests {
		o, _ := BWT(tt.in)
		if o != tt.out {
			t.Errorf("BWT(%q) => %q, expected %q", tt.in, o, tt.out)
		}
	}
}

// var IBWTtests = []struct {
// 	in  string
// 	out string
// }{
// 	{"hye", "hey"},
// 	{"nnbaaa", "banana"},
// }

// func TestIBWT(t *testing.T) {
// 	for _, tt := range IBWTtests {
// 		o := IBWT(tt.in)
// 		if o != tt.out {
// 			t.Errorf("IBWT(%q) => %q, expected %q", tt.in, o, tt.out)
// 		}
// 	}

// }
