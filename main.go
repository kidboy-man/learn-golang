package main

import (
	"fmt"

	"github.com/learn-golang/palindrome"
)

func main() {
	word := "rAcEcar"
	fmt.Printf("%s is a palindrome: %t \n", word, palindrome.IsPalindrome(word))
}
