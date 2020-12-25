package main

import (
	"fmt"
)

func main() {
	var slice1 []float64
	fmt.Println(len(slice1), slice1)

	slice2 := append(slice1, 30)
	slice3 := append(slice2, 60)
	fmt.Println(slice1, slice2, slice3)

	slice3[0] = 40

	fmt.Println(slice1, slice2, slice3)
	fmt.Println()

	for i := 0; i < len(slice3); i++ {
		fmt.Println("slice3:", i, slice3[i])
	}

	fmt.Println()
	fmt.Println("variadic output: ")
	variadic(50, 70, 90, 100)
	fmt.Println()

	slice4 := []string{"a", "b", "c", "d", "e", "f"}
	slice5 := slice4[0:4]
	slice6 := slice4[3:]
	fmt.Println(slice5, slice6)
	slice4[3] = "XX"
	fmt.Println(slice5, slice6)
	slice6[0] = "F"
	fmt.Println(slice5, slice6)

	fmt.Println()
	slice7 := make([]float64, 10)
	fmt.Println(slice7, len(slice7))
	slice7[5] = 7
	fmt.Println(slice7, len(slice7))
	slice8 := append(slice7, 80)
	fmt.Println(slice7, slice8)
}

func variadic(number ...float64) {
	for i, n := range number {
		fmt.Println("index", i, "has value", n)
	}
}
