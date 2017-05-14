package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number:", float64(e))

}

//fmt.Sprint(e) will call e.Error() to convert the value e to a string. If the
////Error() method calls fmt.Sprint(e), then the program recurses until out of
//memory.

//You can break the recursion by converting the e to a value without a String or
////Error method.

func Sqrt(x float64) (z float64, err ErrNegativeSqrt) {
	if x < 0 {
		err = ErrNegativeSqrt(x)

	} else {
		z = 1.0
		var next float64
		for i := 0; i < 10; i++ {
			next = z - (z*z-x)/(2*z)
			z = next

		}

	}
	return z, err

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

}
