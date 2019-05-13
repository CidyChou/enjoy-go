package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	i, j := Rand100i()
	fmt.Println(j)
	fmt.Println(Rand100(i))

	time.Sleep(time.Duration(10000) * time.Millisecond)

	fmt.Println(j)

}

func Rand100(i int64) int {
	rand.Seed(i)
	return rand.Intn(100)
}

func Rand100i() (int64, int) {
	i := time.Now().UnixNano()
	rand.Seed(i)
	return i, rand.Intn(100)
}
