package main

import "fmt"

func cars(x string, z int) (string, int) {
	a := x
	b := z

	return a, b
}

func main() {
	fmt.Println(cars("hi", 5))
}
