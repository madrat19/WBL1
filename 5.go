// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться

package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Print("Specify the running time in seconds: ")
	fmt.Scan(&n)
	ch := make(chan int)
	go send(ch)
	go read(ch)

	// Ждем n секунд
	time.Sleep(time.Second * time.Duration(n))
}

// Отправляет в канал
func send(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

// Читает из канала и выводит в stdout
func read(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
