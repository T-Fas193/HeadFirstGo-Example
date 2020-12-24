package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	// fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)

	successFlag := false

	for counter := 0; counter < 10; counter++ {
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if number > target {
			fmt.Println("Oops, Your guess was HIGH.")
		} else if number < target {
			fmt.Println("Oops, Your guess was LOW. ")
		} else {
			fmt.Println("Good job! You guessed it!")
			successFlag = true
			break
		}
	}

	if successFlag == false {
		fmt.Println("Sorry. You didn't guess my number. It was: [", target, "]")
	}
}
