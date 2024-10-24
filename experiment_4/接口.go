package main

import "fmt"

type Duck interface {
	Speak(s string)
	run()
}


type oneDuck struct {
	name string
	age  int
}

func (one oneDuck) Speak(s string) {
	fmt.Printf("one duck talks:%s\n", s)
}

func (one oneDuck) run() {
	fmt.Println(one.name, " is running")
}

func main() {
	aDuck := oneDuck{"YellowDuck", 12}
	aDuck.Speak("hello")
	aDuck.run()
}
