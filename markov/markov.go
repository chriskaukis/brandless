package markov

import (
	"math/rand"
	"strings"
	"time"
)

type Markov struct {
	prefixLength int
	chain        map[string][]string
}

func New(prefixLength int) *Markov {
	rand.Seed(time.Now().UnixNano())
	return &Markov{
		prefixLength: prefixLength,
		chain:        make(map[string][]string),
	}
}

func (m *Markov) Build(text string) {
	// The first key in our chain is the empty string. This makes it easy
	// and obvious how to choose the first key when generating the message.
	// Especially since there is no way to pick a random key unless we keep
	// track of the keys in a separate list.
	words := strings.Split(text, " ")
	prefix := make([]string, m.prefixLength)
	for i := 0; i < len(words); i++ {
		word := words[i]
		key := strings.Join(prefix, " ")
		m.chain[key] = append(m.chain[key], word)
		copy(prefix, prefix[1:])
		prefix[len(prefix)-1] = word
	}
}

func (m *Markov) Generate(count int) string {
	// The first key will be the empty string key/value.
	prefix := make([]string, m.prefixLength)
	var words []string
	for i := 0; i < count; i++ {
		suffixes := m.chain[strings.Join(prefix, " ")]
		if len(suffixes) == 0 {
			break
		}
		next := suffixes[rand.Intn(len(suffixes))]
		words = append(words, next)
		copy(prefix, prefix[1:])
		prefix[len(prefix)-1] = next
	}
	return strings.Join(words, " ")
}
