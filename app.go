package main

import "fmt"

type Point struct {
	// регистр влияет на область видимости
	X int
	Y int
}

func main() {
	structs()
}

func structs() {
	p1 := Point{
		//при оформлении данных в столбец, последний элемент должен быть с запятой, если в строку то нет.
		X: 1,
		Y: 2,
	}

	fmt.Println(p1.Y)
	p1.Y = p1.X + 4

	fmt.Println(p1.Y)
}
