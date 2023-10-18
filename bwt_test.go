package bwt

import "testing"

func TestNaiveBWT(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in  string
		out string
	}{
		"hey":      {"hey", "yh$e"},
		"banana":   {"banana", "annb$aa"},
		"appellee": {"appellee", "e$elplepa"},
		"dogwood":  {"dogwood", "do$oodwg"},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := BWT(tt.in); got != tt.out {
				t.Errorf("BWT(tt.in) != tt.out, got %v", got)
			}
			if got := InverseBWT(tt.out); got != tt.in {
				t.Errorf("InverseBWT(tt.out) != tt.in, got %v", got)
			}
		})
	}
}
