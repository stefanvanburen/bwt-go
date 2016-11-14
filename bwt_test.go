package BWT

import "testing"

var BWTtests = []struct {
	in  string
	out string
}{
	{"hey", "yh$e"},
	{"banana", "annb$aa"},
	{"appellee", "e$elplepa"},
	{"dogwood", "do$oodwg"},
}

func TestBWT(t *testing.T) {
	for _, tt := range BWTtests {
		o := BWT(tt.in)
		if o != tt.out {
			t.Errorf("BWT(%q) => %q, expected %q", tt.in, o, tt.out)
		}
	}
}

var IBWTtests = []struct {
	in  string
	out string
}{
	{"yh$e", "hey"},
	{"annb$aa", "banana"},
	{"e$elplepa", "appellee"},
	{"do$oodwg", "dogwood"},
}

func TestIBWT(t *testing.T) {
	for _, tt := range IBWTtests {
		z := IBWT(tt.in)
		if z != tt.out {
			t.Errorf("IBWT(%q) => %q, expected %q", tt.in, z, tt.out)
		}
	}
}
