package stringtools

import (
	"strings"
	"unicode"
)

// replaceAll will replace any item under old with the value of new in the string
// Returns a string
func replaceAll(s string, new string, old ...string) string {
	for _,i := range old {
		s = strings.Replace(s, i, new, -1)
	}
	return s
}

// TitledString is a Stringer interface implementation, which replaces the characters - and _ by a space
// and capitalizes the first letter of the string
// Returns a string
type TitledString string

func (str TitledString) String() string {
	s := []rune(replaceAll(string(str), " ", "-", "_"))
	s[0] = unicode.ToUpper(s[0])
	return string(s)
}

//ReplaceBetween will replace occurances of t with st (StartTag) and et (EndTag)
func ReplaceBetween(s *string, t, st, et string) {
	var p int = len(t)
	var useSt bool = true
	str := *s

	for {
		i := strings.Index(str, t)
		if i == -1 {
			*s = str
			break
		}

		if useSt {
			str = str[:i] + st + str[i+p:]
		} else {
			str = str[:i] + et + str[i+p:]
		}
		useSt = !useSt
	}
}

func findBetween(s, st, et string, sp int) (int, int) {
	str := s[sp:]
	lst := len(st)
	i := strings.Index(str, st)
	if i == -1 {
		return -1, -1
	}

	j := strings.Index(str[i+lst:], et)
	if j == -1 {
		return i, -1
	}
	return i + lst, i + j - 1
}

func FindBetween(s, st, et string, sp int) string {
	str := s[sp:]
	start, end := findBetween(str, st, et, 0)
	if start == -1 || end == -1 {
		return ""
	}
	return str[start:end]
}

