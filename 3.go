// Дана последовательность чисел: 2, 4, 6, 8, 10
// Найти сумму их квадратов (2^2+3^2+4^2...) с использованием конкурентных вычислений

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// Конкурентно прибавляет к ответу квадраты чисел из массива
func sumSquares(arr []int) int {
	var ans int
	for _, n := range arr {
		wg.Add(1)
		go addSquare(&ans, n)
	}
	wg.Wait()
	return ans
}

// Прибавляет к ответу квадрат числа
func addSquare(ans *int, n int) {
	defer wg.Done()
	*ans += n * n
}

// Тест
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	ans := sumSquares(arr)
	fmt.Println(ans)
}
