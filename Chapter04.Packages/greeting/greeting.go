package greeting

import "fmt"

// Hello say hello
func Hello() {
	fmt.Println("Hello!")
}

// Hi say Hi
func Hi() {
	fmt.Println("Hi!")
}

// AllGreetings call Hello method and Hi method
func AllGreetings() {
	Hello()
	Hi()
}
