package main

import "fmt"

func greetings(name string) string {
	return "Hello " + name
}

func getSum(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func main() {

	name := "Golang"
	fmt.Println(greetings(name))

	sum := getSum(10, 15, 8)
	fmt.Println(sum)
}