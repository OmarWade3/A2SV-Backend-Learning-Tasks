package main

import (
	"fmt"
	"regexp"
	"strings"
)

// func to count frequncies of the words in a string - ignoring puncutation marks
func freq(input string) map[string]int {
	m := make(map[string]int)
	// compile everything excluding puncutation
	re := regexp.MustCompile("[[:punct:]]+")
	trim := re.ReplaceAllLiteralString(input, "")
	// slice the input by space to get the words
	words := strings.Fields(trim)

	for _, word := range words {
		word := string(word)
		_, prs := m[word]
		if !prs {
			m[word] = 1
			continue
		}
		m[word] += 1
	}
	return m
}

// func to check if a string is a palindrome - ignoring punc and space and case sensitivety
func palindrome(input string) bool {
	re := regexp.MustCompile("[[:punct:]]+")
	trim := re.ReplaceAllLiteralString(input, "")
	lower := strings.ToLower(trim)
	word := []rune(lower)
	l := 0
	r := len(word) - 1
	for l < r {
		lc := string(word[l])
		rc := string(word[r])
		if lc == " " {
			l += 1
		} else if rc == " " {
			r -= 1
		} else if lc == rc {
			l += 1
			r -= 1
		} else {
			return false
		}
	}
	return true
}

func main() {
	// tc example for frequncy func
	s := "test case! * te#st Test"
	n := freq(s)
	fmt.Println("Freq words are: ", n)

	// tc example for palindrome func
	in := "w!o w# "
	fmt.Println("\nis '", in, "' palindrome? ", palindrome(in))
}
