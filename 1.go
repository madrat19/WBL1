// Дана структура Human (с произвольным набором полей и методов)
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования)

package main

import "fmt"

// Родительский тип
type Human struct {
	Name string
	Age  int
}

// Метод 1
func (h Human) talk(s string) {
	fmt.Printf("%s says: '%s'\n", h.Name, s)
}

// Метод 2
func (h Human) work() {
	fmt.Printf("%s is working...\n", h.Name)
}

// Метод 3
func (h Human) introduce() {
	fmt.Printf("Hello, my name is %s, i am %d\n", h.Name, h.Age)
}

// Тип наследник 
type Action struct {
	Human
}

func main() {
	// Тест типа Human
	Alex := Human{Name: "Alex", Age: 25}
	Alex.introduce()
	Alex.talk("Let me talk from my heart")
	Alex.work()

	// Тест типа Action
	someAction := new(Action)
	someAction.Name = "Josh"
	someAction.Age = 31
	someAction.introduce()
	someAction.talk("I am embedded in Action")
	someAction.work()
}
