// Реализовать собственную функцию sleep

package main

import (
	"fmt"
	"time"
)

func sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	time.Sleep()
	fmt.Println("Hello there!")
	fmt.Println("Sleeping 3 seconds...")
	sleep(3)
	fmt.Println("Goodbuy!")
}
