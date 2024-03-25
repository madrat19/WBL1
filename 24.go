// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point
// с инкапсулированными параметрами x,y и конструктором

package main

import (
	"fmt"
	"math"
)

// Структура
type point struct {
	x, y float64
}

// Конструктор
func newPoint(x, y float64) point {
	return point{x: x, y: y}
}

// Нахождения расстояния между точками
func distance(p1, p2 point) float64 {
	return math.Pow(math.Pow((p1.x-p2.x), 2)+math.Pow((p1.y-p2.y), 2), 0.5)
}

// Тест
func main() {
	p1 := newPoint(0, 0)
	p2 := newPoint(3, 4)
	dist := distance(p1, p2)
	fmt.Println(dist)
}
