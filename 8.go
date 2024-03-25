// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0

package main

import (
	"fmt"
)

func setBit(n int64, pos uint, val int) int64 {
	mask := int64(1 << pos)
	if val == 1 {
		return n | mask
	}
	return n &^ mask
}

func main() {
	var num int64
	var pos uint
	var val int

	// Ввод числа
	fmt.Print("Введите число типа int64: ")
	fmt.Scanln(&num)

	// Ввод позиции бита (индексация с 0)
	fmt.Print("Введите позицию бита: ")
	fmt.Scanln(&pos)

	// Ввод значения (0 или 1)
	fmt.Print("Введите значение (0 или 1): ")
	fmt.Scanln(&val)

	// Установка бита
	result := setBit(num, pos, val)
	fmt.Printf("Число после установки %d-го бита в %d: %d\n", pos, val, result)
}
