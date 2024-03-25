// Реализовать конкурентную запись данных в map

package main

import (
	"fmt"
	"sync"
)

// Используем mutex для конкурентной записи в map
// И wg для ожидания завершения всех горутин
var (
	mutex sync.Mutex
	wg    sync.WaitGroup
)

func main() {
	squares := map[int]int{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go mapWrite(squares, i, i*i)
	}
	wg.Wait()
	fmt.Print(squares)
}

// Записывает данные в map
func mapWrite(m map[int]int, k, v int) {
	defer mutex.Unlock()
	defer wg.Done()
	mutex.Lock()
	m[k] = v
}
