// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a := 1
	b := 2
	swap1(&a, &b)
	fmt.Printf("a: %d\nb: %d\n", a, b)
	swap2(&a, &b)
	fmt.Printf("a: %d\nb: %d", a, b)
}

// 1-ый способ
func swap1(a, b *int) {
	*a, *b = *b, *a
}

// 2-ой способ
func swap2(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}
