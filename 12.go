// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество

package main

import "fmt"

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree", "apple", "Bob", "apple"}
	temp := map[string]bool{}
	set := []string{}

	// Записываем все уникальыне строки
	for _, s := range strings {
		temp[s] = true
	}

	// Сохраняем уникальные строки в ответ
	for s := range temp {
		set = append(set, s)
	}

	// Выводим ответ
	fmt.Print(set)
}
