// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a,b, значение которых > 2^20

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Тест сложения
	a := big.NewInt(5000000000000000000)
	b := big.NewInt(6000000000000000000)
	fmt.Printf("%d + %d = %d\n", a, b, bigSum(a, b))

	// Тест вычитания
	x := big.NewInt(7000000000000000000)
	y := big.NewInt(4000000000000000000)
	fmt.Printf("%d - %d = %d\n", x, y, bigSub(x, y))

	// Тест умножения
	i := big.NewInt(4000000000000000000)
	j := big.NewInt(2000000000000000000)
	fmt.Printf("%d * %d = %d\n", i, j, bigMul(i, j))

	// Тест деления
	m := big.NewInt(6000000000000000000)
	l := big.NewInt(3000000000000000000)
	fmt.Printf("%d / %d = %d\n", m, l, bigDiv(m, l))

}

func bigSum(a, b *big.Int) *big.Int {
	sum := new(big.Int)
	return sum.Add(a, b)
}

func bigSub(a, b *big.Int) *big.Int {
	sub := new(big.Int)
	return sub.Sub(a, b)
}

func bigMul(a, b *big.Int) *big.Int {
	mul := new(big.Int)
	return mul.Mul(a, b)
}

func bigDiv(a, b *big.Int) *big.Int {
	div := new(big.Int)
	return div.Div(a, b)
}
