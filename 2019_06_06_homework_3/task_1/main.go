package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
)

func main() {
	words, err := TenPopularWordsFromText(`A hobby is something that people like to do when they are not busy with their usual work and have some free time. It is something being done exclusively for pleasure. Hobbies differ like tastes. If you have chosen a hobby according to your character and taste, you are lucky because your life becomes more interesting. Hobbies are divided into four large classes: doing things, making things, collecting things and learning things. The most popular of all hobby groups is doing things. It includes a wide variety of activities, everything from gardening to travelling and from chess to volleyball. Gardening is one of the oldest man's hobbies. It is known that the English are very keen on gardening and growing flowers, especially roses. Both grown-ups and children are fond of playing different computer games. This is a relatively new hobby but it is becoming more and more popular. Making things includes drawing embroidery, painting, making sculptures, designing costumes, handicrafts, joinery, knitting, book binding, fret work. Two of the most famous amateur painters were President Eisenhower and Sir Winston Churchill. Some write music or play musical instruments. Ex-president of the USA Bill Clinton, for example, plays the saxophone. Almost everyone collects something at some period of his\her life: stamps, coins, matchboxes, books, records, postcards, toys, watches. Some collections have no real value. Others become so large and so valuable that they are housed in museums and galleries. Many world-famous collections started in a small way with one or two items. People with a good deal of money often collect paintings, rare books and other objects of art. Often such private collections are given to museums, libraries and public galleries so that others might take pleasure in seeing them. As far as I am concerned, I have always been fond of collecting stamps. My mother had started collecting stamps long before I was born. She gave me her six albums of stamps as a birthday present when I was twelve. Then I continued collecting stamps myself. It helped me to learn a lot about other countries and other people's traditions, the world's fauna and flora. I used to bring the albums to school and sometimes exchanged stamps with my schoolmates. About a year ago my parents bought me a tiny model car, and since then I decided to collect miniature car models. I am fond of cars and car-races. Now I collect car-models from different countries and historical periods. I also try to find out everything about the new models I buy. I read specialised web-sites and try to keep up with the release of new models on Internet forums. I communicate with other people from different countries about car models, but I have to brush up my English. No matter what kind of hobby a person has, he\she always has the opportunity of learning from it. By reading about the things we are interested in we are adding to what we already know. Learning things can be the most exciting aspect of a hobby.`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(words)
}

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
