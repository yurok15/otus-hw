package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func nextInteger(str string, i int) (bool, int) {
	if i != utf8.RuneCountInString(str)-1 && unicode.IsDigit(rune(str[i+1])) {
		nextDigitValue, _ := strconv.Atoi(string(rune(str[i+1])))
		return true, nextDigitValue
	}
	return false, 0
}

func Unpack(str string) (string, error) {
	var result strings.Builder
	switch {
	// // sting is empty
	case str == "":
		return result.String(), nil

	case unicode.IsLetter(rune(str[0])):
		for i, character := range str {
			// if character is no a Digit
			nextIsDigit, nextValue := nextInteger(str, i)
			switch {
			case unicode.IsLetter(character):
				if nextValue > 0 {
					result.WriteString(strings.Repeat(string(character), nextValue-1))
				} else if nextIsDigit {
					continue
				}
				result.WriteString(string(character))

			case unicode.IsDigit(character):
				if nextIsDigit {
					return result.String(), ErrInvalidString
				}
			default:
				switch {
				case strings.Contains(str, `\\\`):
					result.WriteString(`\`)
					result.WriteString(strings.Split(str, `\\\`)[1])
					return result.String(), nil
				case strings.Contains(str, `\\`):
					repeatNum, _ := strconv.Atoi(strings.Split(str, `\\`)[1])
					result.WriteString(strings.Repeat(`\`, repeatNum))
					return result.String(), nil
				case string(character) == `\`:
					if i != utf8.RuneCountInString(str)-2 && unicode.IsDigit(rune(str[i+2])) {
						secondNextDigitValue, _ := strconv.Atoi(string(rune(str[i+2])))
						result.WriteString(strings.Repeat(string(rune(str[i+1])), secondNextDigitValue))
						return result.String(), nil
					}
					result.WriteString(string(rune(str[i+1])))
				}
			}
		}
	default:
		return result.String(), ErrInvalidString
	}

	return result.String(), nil
}
