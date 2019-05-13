package main

import "fmt"
import "math/rand"
import "time"

func Generate_Randnum() int {
	//rand.Seed(time.Now().Unix())
	fmt.Println(time.Now())

	tm2, _ := time.Parse("2019-05-06 12:05:00.5465", "2019-05-06 12:05:00.5465")

	rand.Seed(tm2.Unix())
	rnd := rand.Intn(100)

	fmt.Printf("rand is %v\n", rnd)

	return rnd
}

func main() {
	Generate_Randnum()
}
