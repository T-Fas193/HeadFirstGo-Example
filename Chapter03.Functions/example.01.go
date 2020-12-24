package main

import "fmt"

func main() {
	// sayHi()
	// repeatLine("I'm Xavier. ", 5)
	// paintNeeded(2.3, 7.1)
	// dozen := double(6)
	// fmt.Println(dozen)
	// fmt.Println(double(4.2))
	fmt.Println(status(60.2))
	fmt.Println(status(40))
}

func sayHi() {
	fmt.Println("Hi!")
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}

func paintNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed.\n", area/10.0)
}

func double(number float64) float64 {
	return number * 2
}

func status(grade float64) string {
	if grade > 60 {
		return "pass"
	}
	return "fail"
}
