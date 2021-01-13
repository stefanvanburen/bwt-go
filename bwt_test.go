package bwt

import (
	"testing"

	"github.com/matryer/is"
)

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
	is := is.New(t)

	for _, tt := range NaiveBWTtests {
		is.Equal(NaiveBWT(tt.in), tt.out)
		is.Equal(NaiveInverseBWT(tt.out), tt.in)
	}
}

var BWTtests = []struct {
	in  string
	out string
}{
	{"hey", "hye"},
	{"banana", "nnbaaa"},
	{"appellee", "eelplepa"},
	{"dogwood", "odoodwg"},
}

func TestBWT(t *testing.T) {
	is := is.New(t)

	for _, tt := range BWTtests {
		o, _ := BWT(tt.in)
		is.Equal(o, tt.out)

		// TODO: this method doesn't work
		// i := IBWT(tt.out, il)
		// assert.Equal(t, tt.in, i)
	}
}
