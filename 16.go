// Реализовать быструю сортировку массива (quicksort) встроенными методами языка

package main

import (
	"fmt"
	"math/rand"
)

func quicksort(arr []int, start, end int) {
	if end-start > 1 {
		p := partition(arr, start, end)
		quicksort(arr, start, p)
		quicksort(arr, p+1, end)
	}
}

func partition(arr []int, start, end int) int {
	i := start + 1
	j := end - 1
	pivot := arr[start]

	for {
		for i <= j && arr[i] <= pivot {
			i++
		}
		for i <= j && arr[j] >= pivot {
			j--
		}

		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
		} else {
			arr[start], arr[j] = arr[j], arr[start]
			return j
		}
	}
}

// Тест
func main() {
	arr := []int{}

	// Заполняем массив 10 случайными числами от -10 до 10
	for i := 0; i < 10; i++ {
		randomNumber := rand.Intn(21) - 10
		arr = append(arr, randomNumber)
	}

	fmt.Println("Исходный массив:", arr)
	quicksort(arr, 0, len(arr))
	fmt.Println("Отсортированный массив:", arr)
}
