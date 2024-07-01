package main

import (
	"fmt"
	"slices"
	"strings"
)

func findAnagram(words *[]string) *map[string]*[]string {
	anagrams := make(map[string]*[]string)
	result := make(map[string]*[]string)

	for _, word := range *words {
		w := []rune(strings.ToLower(word))
		slices.Sort(w)

		if _, ok := anagrams[string(w)]; !ok {
			wordsSet := make([]string, 0)
			anagrams[string(w)] = &wordsSet
		}

		*anagrams[string(w)] = append(*anagrams[string(w)], word)
	}

	for _, v := range anagrams {
		if len(*v) > 1 {
			slices.Sort(*v)
			result[(*v)[0]] = v
		}
	}

	return &result
}

func main() {
	example := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	for k, v := range *findAnagram(&example) {
		fmt.Println(k, ":", *v)
	}
}
