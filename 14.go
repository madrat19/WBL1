// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 1
	b := 3.14
	c := "foo"
	d := make(chan int)
	e := true

	// Тест 1-го способа
	fmt.Println("1-ый способ: ")
	fmt.Print("------------------------------------\n")
	printType(a)
	printType(b)
	printType(c)
	printType(d)
	printType(e)

	// Тест 2-го способа
	fmt.Println("\n2-ой способ: ")
	fmt.Print("------------------------------------\n")
	fmt.Println(getType(a))
	fmt.Println(getType(b))
	fmt.Println(getType(c))
	fmt.Println(getType(d))
	fmt.Println(getType(e))

}

// 1-ый спооб
func printType(i interface{}) {
	fmt.Printf("%T\n", i)
}

// 2-ой способ
func getType(i interface{}) reflect.Type {
	return reflect.TypeOf(i)
}
