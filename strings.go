package tripismapiutilities

import (
	"bytes"
	"fmt"
	"strings"
)

// TitleCase returns title cased string :) all letters that begin words mapped to their title case.
// Certain small words are skipped to ensure grammatical correctness.
// e.g. WELCOME TO THE DOLLHOUSE => Welcome to the Dollhouse
func TitleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	words := strings.Fields(s)
	smallwords := " a an on the to of in and "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = strings.ToLower(word)
		} else {
			words[index] = strings.Title(strings.ToLower(word))
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
