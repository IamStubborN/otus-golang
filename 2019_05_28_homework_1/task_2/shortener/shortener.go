package shortener

import (
	"math/rand"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const base = len(alphabet)
const path = "https://otus.ru/"

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type LinksFormatter struct {
	Links map[string]string
}

func (lf *LinksFormatter) Shorten(url string) string {
	s := make([]rune, 5)
	for i := range s {
		s[i] = []rune(alphabet)[rand.Intn(base)]
	}
	shortenURL := path + string(s)
	lf.Links[shortenURL] = url
	return shortenURL
}

func (lf *LinksFormatter) Resolve(url string) string {
	if resolvedURL, ok := lf.Links[url]; ok {
		return resolvedURL
	}
	return ""
}

func NewLinksFormatter() *LinksFormatter {
	lf := LinksFormatter{}
	lf.Links = make(map[string]string)
	return &lf
}
