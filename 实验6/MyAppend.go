package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}
	slice := []int{}

	slice = append(slice, slice1...)
	slice = append(slice, slice2...)

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice)
}
