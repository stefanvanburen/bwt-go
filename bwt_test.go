package bwt

import (
	"testing"

	"github.com/matryer/is"
)

func TestNaiveBWT(t *testing.T) {
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
		t.Run(name, func(t *testing.T) {
			is := is.New(t)

			is.Equal(BWT(tt.in), tt.out)
			is.Equal(InverseBWT(tt.out), tt.in)
		})
	}
}
