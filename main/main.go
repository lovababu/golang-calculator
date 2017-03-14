package main

import (
	"github.com/lovababu/go-calculator-util/calculator"
	"fmt"
)

func main()  {
	fmt.Println("6 + 4 = " , calculator.Sub(4, 2))
	fmt.Println("6 - 4 = ", calculator.Sub(4, 6))
	fmt.Println("6 * 4 = ", calculator.Multi(6, 4))
	fmt.Println("Sqrt(9) = ", calculator.Sqrt(9))
}
