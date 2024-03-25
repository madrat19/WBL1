// Реализовать все возможные способы остановки выполнения горутины

package main

import (
	"context"
	"fmt"
	"sync"
)

var (
	wg       sync.WaitGroup
	stopFlag bool
)

func main() {
	fmt.Println("-------------------------------------------------------------")
	ch1 := make(chan string)
	wg.Add(1)
	go goroutine1(ch1)
	ch1 <- "Hello"
	ch1 <- "World"
	close(ch1)
	fmt.Println("-------------------------------------------------------------")

	ch2 := make(chan string)
	wg.Add(1)
	go goroutine2(ch2)
	ch2 <- "GO"
	ch2 <- "GO GO GO"
	close(ch2)
	fmt.Println("-------------------------------------------------------------")

	ch3 := make(chan string)
	stop := make(chan bool)
	wg.Add(1)
	go goroutine3(ch3, stop)
	ch3 <- "Good morning"
	ch3 <- "Good night"
	stop <- true
	fmt.Println("-------------------------------------------------------------")

	ch4 := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go goroutine4(ch4, ctx)
	ch4 <- "Ping"
	ch4 <- "Pong"
	cancel()
	fmt.Println("-------------------------------------------------------------")

	ch5 := make(chan string)
	wg.Add(1)
	go goroutine5(ch5)
	ch5 <- "I dont know"
	ch5 <- "What else to send"
	Stop()
	fmt.Println("-------------------------------------------------------------")

	wg.Wait()
}

// Останавливается с помощью закрытия канала
func goroutine1(ch chan string) {
	defer wg.Done()
	fmt.Println("Goroutine 1 started")
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("Goroutine 1 stoped")
			return
		}
		fmt.Println(v)
	}
}

// Почти то же самое, но с использованием цикла range по каналу
func goroutine2(ch chan string) {
	defer wg.Done()
	defer fmt.Println("Goroutine 2 stoped")
	fmt.Println("Goroutine 2 started")
	for v := range ch {
		fmt.Println(v)
	}
}

// Останавливается с помощью дополнительного канала-флага
func goroutine3(ch chan string, stop chan bool) {
	defer wg.Done()
	fmt.Println("Goroutine 3 started")
	for {
		select {
		case <-stop:
			close(ch)
			fmt.Println("Goroutine 3 stoped")
			return
		default:
			fmt.Println(<-ch)
		}
	}
}

// Останавливается синалом отмены, посылаемым с помощью пакаета context
func goroutine4(ch chan string, ctx context.Context) {
	defer wg.Done()
	fmt.Println("Goroutine 4 started")
	for {
		select {
		case <-ctx.Done():
			close(ch)
			fmt.Println("Goroutine 4 stoped")
			return
		default:
			fmt.Println(<-ch)

		}
	}
}

// Остановка с помощью переменной-флага
func goroutine5(ch chan string) {
	defer wg.Done()
	defer fmt.Println("Goroutine 5 stoped")
	fmt.Println("Goroutine 5 started")
	for !shouldStop() {
		fmt.Println(<-ch)
	}
}

func shouldStop() bool {
	return stopFlag
}

func Stop() {
	stopFlag = true
}
