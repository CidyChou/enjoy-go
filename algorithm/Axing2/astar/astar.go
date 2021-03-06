package astar

import (
	"container/heap"
	"container/list"
	"math"
)

// Point is 棋子坐标
type Point struct {
	X int
	Y int
}

// Area is G(x) H(x)
type Area struct {
	Point
	G int
	H int
}

type closeList [][]bool
type openList []Area

func (a Point) getManhattanDistance(b Point) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

func (closeList) New(x, y int) closeList {
	res := make([][]bool, x)
	for i := 0; i < x; i++ {
		res[i] = make([]bool, y)
	}
	return res
}

func (ls openList) Len() int {
	return len(ls)
}

func (ls openList) Less(a, b int) bool {
	return ls[a].G < ls[b].G
}

func (ls openList) Swap(a, b int) {
	ls[a], ls[b] = ls[b], ls[a]
}

func (ls *openList) Push(x interface{}) {
	*ls = append(*ls, x.(Area))
}

func (ls *openList) Pop() interface{} {
	lenth := (*ls).Len()
	res := (*ls)[lenth-1]
	*ls = (*ls)[0 : lenth-1]
	return res
}

func getPath(start, end Point, pathPre map[Point]Point) *list.List {
	pathList := list.New()
	pathCur := end

	for !(pathCur == start) {
		pathList.PushFront(pathCur)
		pathCur = pathPre[pathCur]
	}

	pathList.PushFront(start)
	return pathList
}

// Search is 寻路
func Search(legal [][]int, start, end Point) (int, *list.List) {
	row := len(legal)
	column := len(legal[0])
	var closeList closeList
	closeList = closeList.New(row, column) //将起点区域添加到open列表中
	var openList openList
	openList = make([]Area, row*column)
	pathPre := map[Point]Point{}
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	openList[0] = Area{start, 0, 0}

	for len(openList) > 0 {
		cur := openList.Pop().(Area)
		closeList[cur.X][cur.Y] = true
		if cur.Point == end {
			return cur.G, getPath(start, end, pathPre)
		}

		for _, v := range dir {
			x := cur.X + v[0]
			y := cur.Y + v[1]
			g := cur.G + 1
			h := end.getManhattanDistance(cur.Point)

			if x >= 0 && x < row && y >= 0 && y < column && legal[x][y] != 0 && !closeList[x][y] {

				inopen := false
				tar := Area{Point{x, y}, g, h}

				for index := 0; index < len(openList); index++ {
					if openList[index].Point == tar.Point {
						inopen = true

						if openList[index].G > g {
							openList[index].G = g
							pathPre[tar.Point] = cur.Point
							break
						}
					}
				}

				if !inopen {
					heap.Push(&openList, tar)
					pathPre[tar.Point] = cur.Point
				}
			}
		}
	}

	return -1, nil
}
