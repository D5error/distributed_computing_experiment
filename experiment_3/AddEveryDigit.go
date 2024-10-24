package main

import "fmt"


func digitSum(n int) int {
	fmt.Println("当前数字：", n)
	if n == 0 {
		return 0
	} else {
		return n % 10 + digitSum(n / 10)
	}
}


func main() {
	num := 1234
	fmt.Println("num:", num)
	fmt.Println("digitSum(num):", digitSum(num))
	
}