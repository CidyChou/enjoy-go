package main

import (
	"fmt"
)

func main() {
	// b := new(bytes.Buffer)

	// // b := bytes.NewBufferString("YH")
	// b.Write('Y')

	b := []byte{
		0x89, 0x72,
	}
	fmt.Print(b)
	// 89,72,232,3,1,1,0,0,0,0,0,0,0

}
