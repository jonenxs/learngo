package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("s = %v, len = %d ,cap = %d\n", s, len(s), cap(s))
}

func main() {

	//create slice
	var s []int //Zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)
	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)
	//copy slice
	copy(s2, s1)
	printSlice(s2)
	//delete slice
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	// pop from front
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	// pop from tail
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)

}
