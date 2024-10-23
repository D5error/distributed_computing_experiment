package main

import "fmt"

func main() {
	// myMap := make(map[string] int, 5)
	myMap := map[string] int {
		"apple": 1,
		"banana": 2,
		"orange": 3,
	}
	fmt.Println(myMap);

	fmt.Println(myMap["apple"])

	v, ok := myMap["d5"]
	fmt.Println(v, ok)

	fmt.Println("len:", len(myMap))


	fmt.Println("****myMap:*****")
	for key := range myMap {
		fmt.Println(key, "=", myMap[key])
	}
	fmt.Println("***************")
}