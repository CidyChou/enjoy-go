package main

import (
	"fmt"

	underscore "github.com/ahl5esoft/golang-underscore"
)

func main() {

	var dst []int
	underscore.Range(0, 100, 1).Where(func(r, _ int) bool {
		return r > 20
	}).Value(&dst)

	fmt.Println(dst)
}
