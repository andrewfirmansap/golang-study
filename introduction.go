package main

import "fmt"

func main() {
	var a int = 2
	var b string = "MyWorld"
	c := 100
	fmt.Println(a, "Hello", b, c)
	fmt.Println("Hello, World!")
	// print variable value
	fmt.Printf("%v%v\n", a, b)
	// print variable type
	fmt.Printf("%T %T\n=========\n", a, b)

	var (
		id   string = "ID001"
		name string = "MyName"
	)
	fmt.Printf("userID:%v. userName:%v", id, name)

	var n1 float32 = 10
	var n2 float32 = 55

	fmt.Println(n1 / n2)
	if n1 < n2 {
		fmt.Println("N1 smaller")
	} else {
		fmt.Println("N2 smaller")
	}

	var isN1Bigger bool = false
	n1 = 5000
	isN1Bigger = (n1 > n2)
	if isN1Bigger {
		fmt.Println("N1 bigger")

	} else {
		fmt.Println("N1 smaller")
	}

	// Constant
	const CONSTA int = 10
	const CONSTB float32 = 99.98

	fmt.Printf("CONSTA val: %v\n", CONSTA)
	fmt.Printf("CONSTB val: %v\n", CONSTB)
	// CONSTA = 33
	// CONSTB = 44.1
	// fmt.Printf("CONSTA val: %v\n", CONSTA)
	// fmt.Printf("CONSTB val: %v\n", CONSTB)
	for i := 0; i < 10; i++ {
		fmt.Printf("Normal loop-%v\n", i)
	}

	var arrayItem []int

	arrayItem2 := []int{3, 4, 2}
	arrayItem = append(arrayItem, 3)
	arrayItem = append(arrayItem, 2)
	arrayItem = append(arrayItem, 8)
	for i := 0; i < len(arrayItem); i++ {
		fmt.Printf("idx-%v. Value: %v\n", i, arrayItem[i])

	}
	fmt.Println(arrayItem2)

	fruits_map := map[string]int{
		"Banana": 9,
		"Grape":  88,
	}
	fmt.Println(fruits_map)

	for key, val := range fruits_map {
		fmt.Printf("Map key:%v, map val:%v\n", key, val)
	}

}
