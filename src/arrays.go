package main

import (
	"fmt"
)

const ARRAY_SIZE int = 5

func bin_search(array []int, l int, r int, n int) int {
	if r >= 1 {
		var mid int = l + (r-l)/2
		if array[mid] == n {
			return mid
		} else if array[mid] > n {
			return bin_search(array, l, mid-1, n)
		} else {
			return bin_search(array, mid+1, r, n)
		}
	} else {
		return -1
	}
}

func search(array []int, n int) int {
	return bin_search(array, 0, len(array), n)
}

func main() {
	var a [5]int
	fmt.Println("emp: ", a)

	var b [ARRAY_SIZE]int
	for i := 0; i < ARRAY_SIZE; i++ {
		b[i] = i + 2
	}
	fmt.Println(b)

	for index, val := range b {
		fmt.Printf("index: %d, val: %d\n", index, val)
	}

	var c = make([]int, 100, 100)
	for i := 0; i < 100; i++ {
		c[i] = i
	}

	var index int = search(c, 67)
	fmt.Printf("index: %d == %d\n", index, c[index])
}
