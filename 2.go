// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2, 4, 6, 8, 10)
// и выведет их квадраты в stdout

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// Создает массив для ответа и конкуретно сохраняет в него квадраты чисел
func squares(arr []int) []int {
	ans := make([]int, len(arr))
	for i, n := range arr {
		wg.Add(1)
		go saveSquare(ans, i, n)
	}
	wg.Wait()
	return ans
}

// Записывает квадрат числа в массив
func saveSquare(ans []int, i, n int) {
	defer wg.Done()
	ans[i] = n * n
}

// Тест
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	ans := squares(arr)
	fmt.Println(ans)
}
