package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func makeNumber(n int) (func(int), func() int, func(int), func(int)) {
	number := n
	setter := func(n int) {
		number = n
	}
	getter := func() int {
		return number
	}
	adder := func(n int) {
		number += n
	}
	multiplier := func(n int) {
		number *= n
	}
	return setter, getter, adder, multiplier
}

func main() {
	// pos, neg := adder(), adder()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(
	// 		pos(i),
	// 		neg(-2*i),
	// 	)
	// }
	setter, getter, add, mult := makeNumber(10)
	add(1)
	mult(2)
	fmt.Println(getter()) // 22
	setter(0)
	fmt.Println(getter()) // 0
}
