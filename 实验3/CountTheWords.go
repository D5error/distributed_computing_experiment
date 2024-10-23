package main

import (
	"fmt"
	"strings"
)

func count(str string) map[string]int {
	wordsArr := strings.Split(str, " ")
	retMap := make(map[string]int)

	for word := range wordsArr {
		key := wordsArr[word]
		retMap[key]++
	}
	return retMap	
}


func main() {
	str := "how do you do"
	fmt.Println("str:", str)
	fmt.Println("result:", count(str))
}
