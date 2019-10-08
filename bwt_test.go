package bwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	for _, tt := range NaiveBWTtests {
		assert.Equal(t, NaiveBWT(tt.in), tt.out)
		assert.Equal(t, NaiveInverseBWT(tt.out), tt.in)
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
	for _, tt := range BWTtests {
		o, _ := BWT(tt.in)
		assert.Equal(t, o, tt.out)

		// TODO: this method doesn't work
		// i := IBWT(tt.out, il)
		// assert.Equal(t, tt.in, i)
	}
}
