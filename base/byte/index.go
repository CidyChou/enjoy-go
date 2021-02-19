package main

import "fmt"

func main() {
	// 232,3,1,1,0,0,0,0,0,0,0
	b := make([]byte, 0, 5)
	b = append(b, 'Y')
	b = append(b, 'H')
	b = append(b, 126, 0, 0, 0, 12, 0, 0, 0, 82, 79, 66, 79, 84, 95, 85, 73, 68, 49, 50, 54, 2, 0, 0, 0, 67, 78)

	printSlice(b)
}

func printSlice(x []byte) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
