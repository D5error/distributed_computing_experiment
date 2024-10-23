package main

import "fmt"

type Mover interface {
	Move()
}

type Dog struct {
	name string
}

func (d *Dog) Move() {
	fmt.Println(d.name, "狗正在移动")
}

func main() {
	var mover Mover = &Dog{"黄"}
	fmt.Println("type %T", mover)

	v, ok := mover.()

	if ok {
		fmt.Println(v)
	}
}
