package main

import "fmt"

func main() {

	i := 0
	for i < 6 {
		fmt.Printf("%d, ", i)
		i += 1
	}
	fmt.Printf("\n")

	i = 0
	for j := 5; j > 2; j-- {
		fmt.Printf("%d, ", j)
	}
	fmt.Println(i)
}
