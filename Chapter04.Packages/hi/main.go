package main

import (
	"fmt"
	"log"

	"../greeting"
	"../keyboard"
)

func main() {
	// greeting.Hello()
	greeting.Hi()
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
