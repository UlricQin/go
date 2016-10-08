package str

import (
	"regexp"
	"strings"
)

func IsMatch(s, pattern string) bool {
	match, err := regexp.Match(pattern, []byte(s))
	if err != nil {
		return false
	}

	return match
}

func IsIdentifier(s string, pattern ...string) bool {
	defpattern := "^[a-zA-Z0-9\\-\\_\\.]+$"
	if len(pattern) > 0 {
		defpattern = pattern[0]
	}

	return IsMatch(s, defpattern)
}

func IsMail(s string) bool {
	return IsMatch(s, `\w[-._\w]*@\w[-._\w]*\.\w+`)
}

func IsPhone(s string) bool {
	if strings.HasPrefix(s, "+") {
		return IsMatch(s[1:], `\d{13}`)
	} else {
		return IsMatch(s, `\d{11}`)
	}
}

func Dangerous(s string) bool {
	if strings.Contains(s, "<") {
		return true
	}

	if strings.Contains(s, ">") {
		return true
	}

	if strings.Contains(s, "&") {
		return true
	}

	if strings.Contains(s, "'") {
		return true
	}

	if strings.Contains(s, "\"") {
		return true
	}

	return false
}
