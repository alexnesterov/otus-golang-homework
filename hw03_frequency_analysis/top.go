package hw03frequencyanalysis

import (
	"cmp"
	"regexp"
	"slices"
	"strings"
)

type Word struct {
	Text  string
	Count int
}

type Words []Word

var trimmer = regexp.MustCompile(`^[[:punct:]]+|[[:punct:]]+$`)

func trimWord(word string) string {
	trimmed := trimmer.ReplaceAllString(word, "")
	if trimmed == "" && len(word) > 1 {
		return word
	}
	return trimmed
}

func countWords(fields []string) map[string]int {
	countMap := make(map[string]int)
	for _, word := range fields {
		word = strings.ToLower(word)
		word = trimWord(word)
		if word == "" {
			continue
		}
		countMap[word]++
	}
	return countMap
}

func mapToWords(countMap map[string]int) Words {
	words := make(Words, 0, len(countMap))
	for key, value := range countMap {
		words = append(words, Word{key, value})
	}
	return words
}

func sortWords(words Words) Words {
	result := slices.Clone(words)
	slices.SortFunc(result, func(a, b Word) int {
		if a.Count == b.Count {
			return cmp.Compare(a.Text, b.Text)
		}
		return cmp.Compare(b.Count, a.Count)
	})
	return result
}

func topWords(words Words, num int) []string {
	result := make([]string, 0, num)
	for _, word := range words[:min(num, len(words))] {
		result = append(result, word.Text)
	}
	return result
}

func Top10(s string) []string {
	fields := strings.Fields(s)

	countMap := countWords(fields)
	words := mapToWords(countMap)
	sortedWords := sortWords(words)
	result := topWords(sortedWords, 10)

	return result
}
