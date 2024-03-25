/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

package main

import "fmt"

var justString string

// В примере кода глобальная переменная 'justString' ссылается на локальную переменную 'v' размером 1024 байта
// Таким образом большая и ненужная переменная останется в памяти после завершения работы функции
// Вместо этого мы сделаем копирование во временную переменную 'temp',
// которая освободит память после завершения работы функции
// Также в программе используются срезы рун вместо срезов байт, чтобы мы всегда имели предсказуемое
// количество символов (100) в переменной 'justString' даже при работе с кириллицей и т.д.
func someFunc() {
	v := []rune(createHugeString(1 << 10))
	temp := make([]rune, 100)
	copy(temp, v)
	justString = string(temp)

}

func main() {
	fmt.Println(justString)
	someFunc()
	fmt.Println(justString)
}

// Тестовая реализация
func createHugeString(len int) string {
	return "аааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааа"
}