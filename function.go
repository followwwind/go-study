package main

import "fmt"

func minMax(a int, b int, c int) (int, int){
	max := a
	min := a
	if max < b {
		max = b
	}else{
		min = b
	}

	if max < c {
		max = c
	}else{
		min = c
	}

	return max, min
}

func main() {
	fmt.Println(minMax(1, 2, 3))
}
