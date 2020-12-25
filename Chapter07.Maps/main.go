package main

import "fmt"

func main() {
	fmt.Println("call basic:")
	basic()
	fmt.Println()

	fmt.Println("call loop:")
	loop()
}

func basic() {
	var maps map[string]float64
	fmt.Println(maps)
	maps2 := make(map[string]float64, 3)
	fmt.Println(maps2)
	maps2["my age"] = 181
	fmt.Println(maps2)
	maps3 := map[string]float64{"tallest": 19, "shortest": 8}
	fmt.Println(maps3)
}

func loop() {
	var map1 = map[string]string{"S1": "BLUE", "S8": "CNI", "S10": "KN"}
	for name := range map1 {
		fmt.Println(name)
	}

	for _, value := range map1 {
		fmt.Println(value)
	}

	fmt.Println("getting value from map:", map1["S7"], map1["S1"])

	value, ok := map1["S10"]
	fmt.Println(value, ok)

	value, ok = map1["S9"]
	fmt.Println(value, ok)

	_, ok = map1["S8"]
	fmt.Println(ok)

	delete(map1, "S8")
}
