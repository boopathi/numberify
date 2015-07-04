package numberify

import (
	"regexp"
	"strconv"
	"strings"
)

func Numberify(str string) string {
	words := strings.Fields(str)
	rp := regexp.MustCompile("\\w+")
	tokens := make([]string, 0)
	for _, word := range words {
		word = rp.ReplaceAllStringFunc(word, func(s string) string {
			var (
				fst string
				lst string
				occ int
			)
			if len(s) < 2 {
				return s
			}
			fst = string([]rune(s)[0])
			lst = string([]rune(s)[len(s)-1])
			occ = len(s) - 2
			return fst + strconv.Itoa(occ) + lst
		})
		tokens = append(tokens, word)
	}
	return strings.Join(tokens, " ")
}
