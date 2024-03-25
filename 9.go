// Разработать конвейер чисел
// Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch1 := make(chan int)
	ch2 := make(chan int)

	wg.Add(3)
	go send1(arr, ch1)
	go send2(ch1, ch2)
	go display(ch2)
	// Ждем завершения работы всех горутин
	wg.Wait()

}

// Отправляет из массива в первый канал
func send1(arr []int, ch chan int) {
	defer wg.Done()
	defer close(ch)
	for _, n := range arr {
		ch <- n
	}
}

// Отправляет квадраты чисел из первого канала во второй канал
func send2(in chan int, out chan int) {
	defer wg.Done()
	defer close(out)
	for n := range in {
		out <- n * n
	}
}

// Выводит числа из второго канала в stdout
func display(ch chan int) {
	defer wg.Done()
	for n := range ch {
		fmt.Println(n)
	}
}
