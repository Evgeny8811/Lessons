package main

import "fmt"

func main() {
	i := isTest(2)

	switch {
	case i == 1:
		fmt.Println("равно 1")
	case i == 2:
		fmt.Println("равно 2")
	case i == 3:
		fmt.Println("равно 3")
	default:
		fmt.Println(fmt.Sprintf("равно %x", i))
	}
}

func isTest(x int) int {
	if x == 1 {
		return 1
	} else if x == 2 {
		return 2
	} else if x == 3 {
		return 3
	} else {
		return 4
	}
}
