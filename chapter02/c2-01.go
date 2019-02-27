package main

import "fmt"

func main() {

	i, j := 1, "haha"
	fmt.Println(i, j)
	fmt.Println(mathod(i, i))

}

func mathod(n1 int, n2 int) int {
	return n1 + n2
}
