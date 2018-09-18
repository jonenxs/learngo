package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:]
	fmt.Println("arr[2:] = ", s1)
	s2 := arr[:]
	fmt.Println("arr[:] = ", s2)

	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	//slice 扩展 只可以向后扩展，不可以向前扩展

	slice1 := arr[2:6]
	fmt.Printf("slice1=%v, len(slice1)=%d, cap(slice1)=%d \n", slice1, len(slice1), cap(slice1))
	slice2 := slice1[3:5]
	fmt.Printf("slice2=%v, len(slice2)=%d, cap(slice2)=%d \n", slice2, len(slice2), cap(slice2))

	slice3 := append(slice2, 10)
	slice4 := append(slice3, 11)
	slice5 := append(slice4, 12)

	fmt.Println("slice3, slice4, slice5 = ", slice3, slice4, slice5)
	// slice4 and slice5 no longger
	fmt.Println("arr =", arr)
}
