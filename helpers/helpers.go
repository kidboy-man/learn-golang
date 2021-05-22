package helpers

import (
	"strings"
)

func Reverse(chars string) string {
	sliceOfChars := strings.Split(chars, "")
	lenSlice := len(sliceOfChars)
	var sliceOfReversed []string
	for i := range sliceOfChars {
		sliceOfReversed = append(sliceOfReversed, sliceOfChars[(lenSlice-1)-i])
	}
	return strings.Join(sliceOfReversed, "")
}
