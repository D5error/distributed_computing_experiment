package main

import "fmt"

type Mover interface {
	Move()
}

type Doger interface {
	Move()
}

type Dog struct {
	name string
}

func (d Dog) Move() {
	fmt.Println(d.name, "狗正在移动")
}

func main() {
	var mover Mover
	var dog1 = Dog{"黄"}
	mover = dog1

	mover.Move()

	var dog2 = &Dog{"白"}
	mover = dog2

	mover.Move()
}
