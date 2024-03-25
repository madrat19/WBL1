// Реализовать пересечение двух неупорядоченных множеств

package main

import "fmt"

func main() {
	// Пример входных данных (Ожидается уникальность элементов в множествах)
	set1 := []int{1, 2, 3, 4, 5, 6}
	set2 := []int{5, 6, 7, 8, 9, 10}

	intersection := []int{}
	// Словарь: {Элемет : Количество}
	temp := map[int]int{}

	// Записываем в словарь количество вхождений каждого элемента в оба множества вместе
	for _, e := range set1 {
		temp[e] += 1
	}

	for _, e := range set2 {
		temp[e] += 1
	}

	// Добавляем в ответ те жлементы, которые встретились по 2 раза
	for e, count := range temp {
		if count == 2 {
			intersection = append(intersection, e)
		}
	}

	// Выводим ответ
	fmt.Print(intersection)
}
