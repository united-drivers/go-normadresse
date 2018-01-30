package normadresse

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type rule struct {
	step  float32
	short string
	long  string
}

func Normalize(input string) string {
	return NormalizeLength(input, 32)
}

func NormalizeLength(str string, maxLen int) string {
	str = removeAccents(strings.ToUpper(str))
	for _, rule := range rules {
		re := regexp.MustCompile(rule.long)
		str = re.ReplaceAllString(str, rule.short)
		if len(str) <= maxLen {
			return str
		}
	}
	str = strings.Replace(str, " @", "", -1)
	return str
}

func removeAccents(input string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ := transform.String(t, input)
	return s
}
