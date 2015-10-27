package lib

import (
	"strings"
)

type Filter struct {
	words []string
}

func NewFilter(words []string) *Filter {
	return &Filter{words: words}
}

func (self *Filter) Filter(target string) (ret map[int]string) {
	ret = make(map[int]string)
	for _, word := range self.words {
		n := strings.LastIndex(target, word)
		for n != -1 {
			ret[n] = word
			n = strings.LastIndex(string(target[0:n]), word)
		}
	}
	return
}