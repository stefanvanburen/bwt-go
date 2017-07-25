package BWT

import "testing"

var NaiveBWTtests = []struct {
	in  string
	out string
}{
	{"hey", "yh$e"},
	{"banana", "annb$aa"},
	{"appellee", "e$elplepa"},
	{"dogwood", "do$oodwg"},
}

func TestNaiveBWT(t *testing.T) {
	for _, tt := range NaiveBWTtests {
		o := NaiveBWT(tt.in)
		if o != tt.out {
			t.Errorf("NaiveBWT(%q) => %q, expected %q", tt.in, o, tt.out)
		}
		i := NaiveIBWT(tt.out)
		if i != tt.in {
			t.Errorf("NaiveIBWT(%q) => %q, expected %q", tt.out, i, tt.in)
		}
	}
}

var BWTtests = []struct {
	in  string
	out string
}{
	{"hey", "hye"},
	// {"banana", "nnbaaa"},
	// {"appellee", "eelplepa"},
	// {"dogwood", "odoodwg"},
}

func TestBWT(t *testing.T) {
	for _, tt := range BWTtests {
		o, il := BWT(tt.in)
		if o != tt.out {
			t.Errorf("BWT(%q) => %q, expected %q", tt.in, o, tt.out)
		}
		i := IBWT(tt.out, il)
		if i != tt.in {
			t.Errorf("IBWT(%q) => %q, expected %q", tt.out, i, tt.in)
		}
	}
}
