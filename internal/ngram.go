package internal

import (
	"sort"
	"strings"
	"unicode"
)

func tokenize(text string) []string {
	text = strings.ToLower(text)
	var cleanedText strings.Builder
	for _, char := range text {
		if unicode.IsLetter(char) || unicode.IsSpace(char) {
			cleanedText.WriteRune(char)
		} else {
			cleanedText.WriteRune(' ')
		}
	}

	return strings.Fields(cleanedText.String())
}

func generateNGrams(tokens []string, n int) []string {
	var ngrams []string
	if len(tokens) < n {
		return ngrams
	}

	for i := 0; i <= len(tokens)-n; i++ {
		ngram := strings.Join(tokens[i:i+n], " ")
		ngrams = append(ngrams, ngram)
	}
	return ngrams
}

func countNGrams(ngrams []string) map[string]int {
	frequencyMap := make(map[string]int)
	for _, ngram := range ngrams {
		frequencyMap[ngram]++
	}
	return frequencyMap
}

func filterTopNGrams(frequencyMap map[string]int, threshold int, k int) []string {
	filteredMap := make(map[string]int)
	for ngram, count := range frequencyMap {
		if count >= threshold {
			filteredMap[ngram] = count
		}
	}

	type ngramCount struct {
		ngram string
		count int
	}

	ngramCounts := make([]ngramCount, 0, len(filteredMap))
	for ngram, count := range filteredMap {
		ngramCounts = append(ngramCounts, ngramCount{ngram, count})
	}

	sort.Slice(ngramCounts, func(i, j int) bool {
		return ngramCounts[i].count > ngramCounts[j].count
	})

	result := []string{}
	for i := 0; i < k && i < len(ngramCounts); i++ {
		result = append(result, ngramCounts[i].ngram)
	}

	return result
}

func FindNGrams(input string) []string {
	tokens := tokenize(input)

	// look exclusively for unigrams and bigrams
	ngrams := generateNGrams(tokens, 2)
	ngrams = append(ngrams, generateNGrams(tokens, 1)...)

	frequencyMap := countNGrams(ngrams)

	return filterTopNGrams(frequencyMap, 5, 3)
}
