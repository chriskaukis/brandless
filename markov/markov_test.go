package markov

import (
	"testing"
)

func TestBuild(t *testing.T) {
	tests := []struct {
		prefixLength int
		input        string
		chain        map[string][]string
	}{
		{
			1,
			"This is a markov chain",
			map[string][]string{
				"":       {"This"},
				"This":   {"is"},
				"is":     {"a"},
				"a":      {"markov"},
				"markov": {"chain."},
			},
		},
		{
			2,
			"This is a markov chain",
			map[string][]string{
				" ":        {"This"},
				" This":    {"is"},
				"This is":  {"a"},
				"is a":     {"markov"},
				"a markov": {"chain."},
			},
		},
	}

	for _, test := range tests {
		markov := New(test.prefixLength)
		markov.Build(test.input)

		if len(markov.chain) != len(test.chain) {
			t.Errorf("got chain of length %d, wanted %d", len(markov.chain), len(test.chain))
		}

		for k := range test.chain {
			if _, ok := markov.chain[k]; !ok {
				t.Errorf("chain missing key %s", k)
			}
		}
	}
}
