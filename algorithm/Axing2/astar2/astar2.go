package astar2

import "container/list"

type Point struct {
	X int
	Y int
}

type Area struct {
	Point
	G int
	H int
}

type closeList [][]bool
type openList []Area

func (ls openList) Len() int {
	return len(ls)
}

func (ls *openList) Pop() interface{} {
	lenth := (*ls).Len()
	res := 
}

// Search2 is 寻路
func Search2(legal [][]int, start, end Point) (int, *list.List) {
	row := len(legal)
	column := len(legal[0])
	var closeList closeList
	closeList = closeList.New(row, column)
	var openList openList
	openList = make([]Area, row*column)
	//pathPre := map[Point]Point{}
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	openList[0] = Area{start, 0, 0}

	for len(openList) > 0 {
		cur := openList.Pop().(Area)
	}

	return -1, nil
}
