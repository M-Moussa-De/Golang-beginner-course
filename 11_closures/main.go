package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}


func main() {

	test := func(x int) int {
		return x * -1
	}(5)


	fmt.Println(test)

	sum := adder()

	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}


}