// Разработать программу, которая переворачивает подаваемую на ход строку
// Например: «главрыба — абырвалг»
// Символы могут быть unicode

package main

import "fmt"

// Тест
func main() {
	s := "а роза упала на лапу Азора"
	fmt.Print(reverse(s))
}

func reverse(s string) string {
	temp := []rune(s)
	i := 0
	j := len(temp) - 1
	for i < j {
		temp[i], temp[j] = temp[j], temp[i]
		i++
		j--
	}
	return string(temp)
}
