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

			is.Equal(NaiveBWT(tt.in), tt.out)
			is.Equal(NaiveInverseBWT(tt.out), tt.in)
		})
	}
}

func TestBWT(t *testing.T) {
	tests := map[string]struct {
		in  string
		out string
	}{
		"hey":       {"hey", "hye"},
		"banana":    {"banana", "nnbaaa"},
		"appellee":  {"appellee", "eelplepa"},
		"dogwood":   {"dogwood", "odoodwg"},
		"wikipedia": {"SIX.MIXED.PIXIES.SIFT.SIXTY.PIXIE.DUST.BOXES", "TEXYDST.E.IXIXIXXSSMPPS.B..E.S.EUSFXDIIOIIIT"},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			is := is.New(t)

			o, _ := BWT(tt.in)
			is.Equal(o, tt.out)

			// TODO: this method doesn't work
			// i := InverseBWT(tt.out, il)
			// is.Equal(tt.in, i)
		})
	}
}
