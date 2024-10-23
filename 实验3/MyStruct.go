package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var point Point
	point.x = 1
	point.y = 2
	fmt.Println("point", point)

	fmt.Println("**** struct pointer ****")
	pointPointer := &point
	fmt.Println("pointPointer:", pointPointer)
	fmt.Println("*pointPointer:", *pointPointer)
	fmt.Println("&pointPointer:", &pointPointer)
	fmt.Println("************************")
}
