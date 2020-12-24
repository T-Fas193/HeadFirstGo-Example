package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var beefOrder = []int{1, 2, 3, 4, 5}
	//for i := 0; i < len(beefOrder); i++ {
	//	fmt.Println(beefOrder[i])
	//}

	for i, data := range beefOrder {
		fmt.Println("index", i, "has value", data)
	}

	fmt.Println("order total is", sumOrder(beefOrder))

	fmt.Println("order total from file is", sumOrderFromFile())
}

func sumOrder(order []int) int {
	var sum int
	for _, data := range order {
		sum += data
	}
	return sum
}

func sumOrderFromFile() float64 {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	if path != "" {
		path = path + "\\resource\\data.txt"
	}
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(file)
	var sum float64
	for scanner.Scan() {
		weight := scanner.Text()
		fmt.Println(weight)
		weight = strings.TrimSpace(weight)
		fweight, err := strconv.ParseFloat(weight, 64)
		if err != nil {
			log.Fatalln(err)
		}
		sum += fweight
	}
	err = file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	if scanner.Err() != nil {
		log.Fatalln(scanner.Err())
	}
	return sum
}
