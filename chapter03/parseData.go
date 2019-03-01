package main

import "fmt"

func main() {
	fmt.Println("开始")

	bytes := []byte{250, 251, 03} //1C 49 05 FF 00 00 00 00 09 00 3E BA 01 00 00 50 28 15 27 9D 59 62 41 01 01 00 0C 46
	arrLen := len(bytes)
	fmt.Println("Data:", bytes, "Len", arrLen)
	// sum := 0

	for i, byte := range bytes {
		fmt.Println("byte", byte)
		if bytes[i] == 250 && bytes[i+1] == 251 {
			fmt.Println("前导码正常")
		}
	}
}
