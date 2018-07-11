package values

import "strings"

type Word struct {
	Lang  string
	Value string
}

func NewWord(lang, value string) Word {
	return Word{
		Lang:  strings.ToLower(lang),
		Value: strings.ToLower(value),
	}
}
