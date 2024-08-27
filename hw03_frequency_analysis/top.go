package hw03frequencyanalysis

import (
	"fmt"
	"sort"
	"strings"
)

type Pair struct {
	K string
	V int
}

func Top10(s string) []string {
	if len(s) == 0 {
		return nil
	}

	// заполняем мэпку с частотами
	freq := make(map[string]int)
	words := strings.Fields(s)
	for _, w := range words {
		freq[w] += 1
	}
	fmt.Println(freq)

	// конвертируем в слайс пар
	pairs := make([]Pair, 0, len(freq))
	for k, v := range freq {
		pairs = append(pairs, Pair{k, v})
	}

	// сортируем слайс пар
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].V == pairs[j].V {
			return pairs[i].K < pairs[j].K
		} else {
			return pairs[i].V > pairs[j].V
		}
	})

	// получаем слайс ключей
	keys := make([]string, 0, len(pairs))
	i := 0
	for i < 10 {
		keys = append(keys, pairs[i].K)
		i++
	}

	return keys
}
