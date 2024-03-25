// Реализовать постоянную запись данных в канал (главный поток)
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout
// Необходима возможность выбора количества воркеров при старте
// Программа должна завершаться по нажатию Ctrl+C.
// Выбрать и обосновать способ завершения работы всех воркеров

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wg sync.WaitGroup

func main() {

	var n int
	fmt.Print("Specify the number of workers: ")
	fmt.Scan(&n)
	wg.Add(n)
	// Добавляем количество воркекров в группу ожидания
	ch := make(chan int)

	go write(ch)

	for i := 1; i <= n; i++ {
		go worker(ch, i)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	// Закрываем канал, чтобы прекратить работу воркеров
	close(ch)

	// Ждем завершения работы всех воркеров
	wg.Wait()
}

// Отправка в канал
func write(ch chan int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ptogram finished")
			os.Exit(0)
		}
	}()
	for i := 1; ; i++ {
		ch <- i
	}
}

// Чтение из канала и вывод в stdout
func worker(ch chan int, n int) {
	for i := range ch {
		fmt.Printf("From worker %v: %v\n", n, i)
	}
}
