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
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := BWT(tt.in, '$')
			if err != nil {
				t.Fatalf("got err %v, expected nil", err)
			}
			if got != tt.out {
				t.Errorf("BWT(tt.in) != tt.out, got %v", got)
			}
			got, err = InverseBWT(tt.out, '$')
			if err != nil {
				t.Fatalf("got err %v, expected nil", err)
			}
			if got != tt.in {
				t.Errorf("InverseBWT(tt.out) != tt.in, got %v", got)
			}
		})
	}
}
