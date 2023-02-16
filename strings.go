package tripismapiutilities

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/runes"
	"golang.org/x/text/secure/precis"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Tokenize simply turns an input string into it's lower case equivalent
func Tokenize(s string) string {
	s = strings.TrimSpace(s)
	return strings.ToLower(s)
}

// TokenizeAndRemoveSpaces turns an input string into it's lower case equivalent and removes any spaces from within the string
func TokenizeAndRemoveSpaces(s string) string {
	if s == "" {
		return s
	}
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "", -1)
	return s
}

// TokenizeAndNormalize turns an input string into it's lower case equivalent then normalizes the text, appending the normalized version to the existing words (split on space) where it differs
// For example, for an input of "Hôtel Longemalle Genève", the returned string will be "hôtel hotel longemalle genève geneve"
func TokenizeAndNormalize(s string) string {
	s = Tokenize(s)
	a := strings.Split(s, " ")
	r := []string{}
	for _, word := range a {
		r = append(r, word)
		norm := RemoveAccents(word)
		if norm != word && len(norm) > 0 {
			r = append(r, norm)
		}
	}
	return strings.Join(r, " ")
}

// RemoveAccents remove diacritics from strings to normalize - useful for tokenization of hotel names to remove accents
func RemoveAccents(s string) string {
	s = replaceChars(s)
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), runes.Remove(runes.In(unicode.P)), norm.NFC)
	output, _, e := transform.String(t, s)
	if e != nil {
		return s
	}
	o := removeAdditionalDiacritics(output)
	return o
}

// removeAdditionalDiacritics is an internal function which handles Polish characters as well as others not caught in standard normalization library
// taken from https://stackoverflow.com/questions/26722450/remove-diacritics-using-go
func removeAdditionalDiacritics(s string) string {
	t := transform.Chain(
		norm.NFD,
		precis.UsernameCaseMapped.NewTransformer(),
		runes.Map(func(r rune) rune {
			switch r {
			// case 'ą':
			// 	return 'a'
			// case 'ć':
			// 	return 'c'
			// case 'ę':
			// 	return 'e'
			case 'ł':
				return 'l'
			// case 'ń':
			// 	return 'n'
			// case 'ó':
			// 	return 'o'
			// case 'ś':
			// 	return 's'
			// case 'ż':
			// 	return 'z'
			// case 'ź':
			// 	return 'z'
			case 'đ':
				return 'd'
			}
			return r
		}),
		norm.NFC,
	)
	output, _, e := transform.String(t, s)
	if e != nil {
		return s
	}
	return output
}

// replaceChars replaces other diacritics which have no available singular rune conversion
// for example, for an input of "ß", the returned string will be "ss"
func replaceChars(s string) string {
	s = strings.ReplaceAll(s, "ß", "ss")
	s = strings.ReplaceAll(s, "&", "and")
	return s
}

// Title is a shortcut method to calling TitleCase
func Title(s string) string {
	return TitleCase(s)
}

// TitleCase returns title cased string :) all letters that begin words mapped to their title case.
// Certain small words are skipped to ensure grammatical correctness.
// e.g. WELCOME TO THE DOLLHOUSE => Welcome to the Dollhouse
func TitleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	words := strings.Fields(s)
	smallwords := " a an on the to of in and "

	caser := cases.Title(language.Und)

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = strings.ToLower(word)
		} else {
			words[index] = caser.String(strings.ToLower(word))
		}
	}
	return strings.Join(words, " ")
}

// LeftCharacters returns only the specified number of characters starting from the first letter of the provided string.
func LeftCharacters(s string, i int) string {
	if len(s) == 0 {
		return s
	}
	if len(s) < i {
		return s
	}
	return s[:i]
}

// RemoveLeftCharacters returns the remainder of a string after the first letters specified have been removed
func RemoveLeftCharacters(s string, i int) string {
	if len(s) == 0 {
		return s
	}
	if len(s) < i {
		return s
	}
	return s[i:]
}

// Truncate truncates a string to the first count of characters specified in i
func Truncate(s string, i int) string {
	if len(s) == 0 {
		return s
	}
	if len(s) < i {
		return s
	}
	return s[:i]
}

// TruncateToWholeWord truncates a string to the first word boundary denoted by a space less than the length specified in i
func TruncateToWholeWord(s string, i int) string {
	if len(s) == 0 {
		return s
	}
	if len(s) < i {
		return s
	}
	var b strings.Builder
	sArray := strings.Split(s[:i], " ")
	for i, s = range sArray {
		//Ignore the last truncated word
		if i == len(sArray)-1 {
			break
		}
		b.WriteString(s)
		b.WriteString(" ")
	}
	return b.String()
}

// ConvertMapToString creates formatted string with key-value pairs from map[string]string
func ConvertMapToString(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\",\n", key, value)
	}
	return b.String()
}
