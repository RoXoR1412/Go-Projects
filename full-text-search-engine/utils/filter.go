package utils

import (
	"strings"

	snowballeng "github.com/kljensen/snowball/english"
)

func lowercaseFilter(token []string) []string {
	r := make([]string, len(token))
	for i, t := range token {
		r[i] = strings.ToLower(t)
	}
	return r
}

func stopwordFilter(tokens []string) []string {
	var stopwords = map[string]struct{}{
		"a":     {},
		"an":    {},
		"the":   {},
		"and":   {},
		"or":    {},
		"but":   {},
		"to":    {},
		"of":    {},
		"in":    {},
		"for":   {},
		"on":    {},
		"with":  {},
		"at":    {},
		"by":    {},
		"this":  {},
		"that":  {},
		"is":    {},
		"was":   {},
		"it":    {},
		"as":    {},
		"be":    {},
		"are":   {},
		"not":   {},
		"from":  {},
		"which": {},
		"all":   {},
		"there": {},
		"their": {},
		"they":  {},
		"we":    {},
		"you":   {},
		"he":    {},
		"she":   {},
		"him":   {},
		"her":   {},
	}

	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := stopwords[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}

func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = snowballeng.Stem(token, false)
	}
	return r
}
