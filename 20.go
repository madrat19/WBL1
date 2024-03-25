// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow»

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Коту тащат уток"
	fmt.Print(reveseWords(s))
}

func reveseWords(s string) string {
	words := strings.Split(s, " ")
	reverseSlice(words)
	ans := strings.Join(words, " ")
	return ans
}

func reverseSlice(slice []string) {
	for i := 0; i < len(slice)/2; i++ {
		j := len(slice) - i - 1
		slice[i], slice[j] = slice[j], slice[i]
	}
}
