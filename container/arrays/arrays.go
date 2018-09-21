package main

import "fmt"

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var array1 [5]int
	array2 := [3]int{1, 2, 3}
	array3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(array1, array2, array3)
	var grid [4][5]int
	fmt.Println(grid)

	for i := 0; i < len(array3); i++ {
		fmt.Println(array3[i])
	}

	for i, v := range array3 {
		fmt.Println(i, v)
	}

	printArray(array1)
}
