package palindrome

import (
	"strings"

	"github.com/learn-golang/helpers"
)

func IsPalindrome(word string) bool {
	cleanedWord := strings.ToLower(strings.TrimSpace(word))
	reversedWord := helpers.Reverse(cleanedWord)
	return cleanedWord == reversedWord
}
