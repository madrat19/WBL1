// Удалить i-ый элемент из слайса

package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("Slice at start: ", arr)

	arr = removeIndex1(arr, 3)
	fmt.Println("Slice after index 3 removed: ", arr)

	arr = removeIndex2(arr, 1)
	fmt.Println("Slice after index 1 removed: ", arr)
}

// 1-ый способ
func removeIndex1(arr []int, index int) []int {
	if index < 0 || index >= len(arr) {
		return arr
	}

	return append(arr[:index], arr[index+1:]...)
}

// 2-ой способ
func removeIndex2(arr []int, i int) []int {
	if i < 0 || i >= len(arr) {
		return arr
	}
	copy(arr[i:], arr[i+1:])
	return arr[:len(arr)-1]
}
