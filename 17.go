// Реализовать бинарный поиск встроенными методами языка

package main

import "fmt"

// Тест
func main() {
	arr := []int{-10, -5, -2, 3, 6, 8, 13, 21}
	x := 8
	pos := binSearch(x, arr)
	if pos == -1 {
		fmt.Printf("Slice does not contain '%d'", x)
	} else {
		fmt.Printf("Found '%d' on position %d", x, pos)
	}

}

func binSearch(x int, arr []int) int {
	l := 0
	r := len(arr) - 1
	for l < r {
		m := (l + r + 1) / 2
		if arr[m] <= x {
			l = m
		} else {
			r = m - 1
		}
	}
	if arr[l] != x {
		return -1
	}
	return l
}
