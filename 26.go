// Разработать программу, которая проверяет, что все символы в строке уникальные
// true — если уникальные, false etc
// Функция проверки должна быть регистронезависимой

package main

import (
	"fmt"
	"strings"
)

// Тест
func main() {
	s1 := "qwerty"
	s2 := "abcdA"
	s3 := "абвгд"
	s4 := "маМА"
	fmt.Println(s1, ":", isUnique(s1))
	fmt.Println(s2, ":", isUnique(s2))
	fmt.Println(s3, ":", isUnique(s3))
	fmt.Println(s4, ":", isUnique(s4))

}

// Подсчитываем количество вхождений каждого символа встроку c помощью map
// И, когда встречается второе вхождение, возвращаем false
// Если второе вхождение ни разу не встретилось, возвразаем true
// Для регистронезависимости используем ToLower()
func isUnique(s string) bool {
	counter := map[rune]int{}
	for _, r := range []rune(strings.ToLower(s)) {
		counter[r]++
		if counter[r] == 2 {
			return false
		}
	}
	return true
}
