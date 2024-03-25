// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде
// По завершению программа должна выводить итоговое значение счетчика

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// Тип-счетчик
type counter struct {
	n     int
	mutex sync.Mutex
}

// Метод для инкремента счетчика
func (c *counter) increment() {
	defer wg.Done()
	defer c.mutex.Unlock()
	c.mutex.Lock()
	c.n++
}

// Тест
func main() {
	var c counter

	// Запускаем 10 горутин
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go c.increment()
	}

	// Ждем завершения всех горутин
	wg.Wait()
	// Выводим значение счетчика
	fmt.Println(c.n)
}
