package main

import (
	"enjoy-go/algorithm/Axing2/astar"

	"fmt"
)

func main() {

	legal := [][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}
	start := astar.Point{0, 0}
	end := astar.Point{3, 3}
	lenth, path := astar.Search(legal, start, end)

	fmt.Println(lenth)

	for p := path.Front(); p != nil; p = p.Next() {
		fmt.Println(p.Value)
	}
}
