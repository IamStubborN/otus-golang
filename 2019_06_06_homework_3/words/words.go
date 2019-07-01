package words

import (
	"regexp"
	"sort"
	"strings"
)

func TenPopularWordsFromText(text string) ([]string, error) {
	textRefactored := removeSeletectedChars(text, `\.|,|'|:|\\|/|"|\d`)
	list := calculatePopularWords(textRefactored)
	return getTenPopularWords(list), nil
}

func removeSeletectedChars(text, regex string) string {
	reg := regexp.MustCompile(regex)
	return reg.ReplaceAllString(text, "")
}

func calculatePopularWords(text string) map[string]int {
	result := make(map[string]int)
	words := strings.Split(text, " ")
	for _, val := range words {
		if len(val) > 0 {
			result[val]++
		}
	}
	return result
}

func getTenPopularWords(list map[string]int) []string {

	type word struct {
		str   string
		count int
	}

	result := make([]string, 0, 10)
	words := make([]word, 0, len(list))
	for key, value := range list {
		words = append(words, word{key, value})
	}
	sort.Slice(words, func(i, j int) bool {
		return words[i].count > words[j].count
	})

	for _, word := range words[0:10] {
		result = append(result, word.str)
	}
	return result
}
