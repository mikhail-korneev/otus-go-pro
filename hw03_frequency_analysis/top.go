package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type WordCountPair struct {
	word  string
	count int
}

func Top10(s string) []string {
	if len(s) == 0 {
		return nil
	}

	freq := make(map[string]int)
	words := strings.Fields(s)
	for _, w := range words {
		freq[w]++
	}

	pairs := make([]WordCountPair, 0, len(freq))
	for word, count := range freq {
		pairs = append(pairs, WordCountPair{word, count})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count == pairs[j].count {
			return pairs[i].word < pairs[j].word
		}
		return pairs[i].count > pairs[j].count
	})

	keys := make([]string, 0, len(pairs))
	top := min(10, len(pairs))
	for i := 0; i < top; i++ {
		keys = append(keys, pairs[i].word)
	}

	return keys
}
