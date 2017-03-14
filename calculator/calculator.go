package calculator

import (
	"math"
)

func Sum(num1, num2 int32) int32 {
	return num1 + num2
}

func Sub(num1, num2 int32) int32 {
	return num2 - num1
}

func Multi(num1, num2 int32) int32 {
	return num1 * num2;
}

func Sqrt(num float64) float64 {
	return math.Sqrt(num)
}
