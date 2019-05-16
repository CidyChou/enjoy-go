package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

var chessboard [][]string
var element []string
var targetPoint map[point]map[point][]point

type point struct {
	X       int
	Y       int
	element string
}

func main() {
	// chessDrop()
	checkPos()

	for c := range chessboard {
		fmt.Println(chessboard[c])
	}

	onTrigger()

	for {
		var p1, p2 string
		fmt.Println("请输入要移动的棋子的位置 与 移动后的位置: eg:(1,1 1,2)")
		fmt.Scanln(&p1, &p2)
		X1, _ := strconv.ParseFloat(strings.Split(p1, ",")[0], 64)
		Y1, _ := strconv.ParseFloat(strings.Split(p1, ",")[1], 64)

		X2, _ := strconv.ParseFloat(strings.Split(p2, ",")[0], 64)
		Y2, _ := strconv.ParseFloat(strings.Split(p2, ",")[1], 64)

		if math.Abs(X2-X1) > 1 || math.Abs(Y2-Y1) > 1 {
			fmt.Println("请输入正确的位置")
		} else {
			p1 := point{int(X1), int(Y1), chessboard[int(X1)][int(Y1)]}
			p2 := point{int(X2), int(Y2), chessboard[int(X2)][int(Y2)]}
			tp1 := targetPoint[p1][p2]
			tp2 := targetPoint[p2][p1] //TODO 优化
			fmt.Println("tp1", tp1, "tp2", tp2)

			if len(tp1) != 0 || len(tp2) != 0 {

				chessboard[int(X1)][int(Y1)], chessboard[int(X2)][int(Y2)] = chessboard[int(X2)][int(Y2)], chessboard[int(X1)][int(Y1)]

				checkPos()
				fmt.Println(" 消除后的棋盘 ")
				for c := range chessboard {
					fmt.Println(chessboard[c])
				}
				onTrigger()
			}
		}
	}
}

func init() {
	chessboard = [][]string{
		{"X", "X", "X", "X", "X", "X", "X", "X"},
		{"X", "X", "X", "X", "X", "X", "X", "X"},
		{"X", "X", "X", "X", "X", "X", "X", "X"},
		{"X", "X", "X", "X", "X", "X", "X", "X"},
		{"X", "X", "X", "X", "X", "X", "X", "X"},
		{"X", "X", "X", "X", "X", "X", "X", "X"},
		{"X", "X", "X", "X", "X", "X", "X", "X"},
		{"X", "X", "X", "X", "X", "X", "X", "X"},
	}

	//targetPoint = make(map[point]point)
	targetPoint = make(map[point]map[point][]point)

	element = []string{"H", "%", "&", "Q", "M"}

	initChessboard(chessboard)

	for c := range chessboard {
		fmt.Println(chessboard[c])
	}

}

func filter(a, b, c point) point {
	if a.element == b.element {
		return c
	}
	if a.element == c.element {
		return b
	}
	if c.element == b.element {
		return a
	}
	return a
}

// 触发器 生成可触发的事件
func onTrigger() {
	for row, cb := range chessboard {
		for col := range cb {
			p := point{row, col, chessboard[row][col]}
			if p.X < len(chessboard)-2 {
				a, b, c := p, point{p.X + 1, p.Y, chessboard[p.X+1][p.Y]}, point{p.X + 2, p.Y, chessboard[p.X+2][p.Y]}
				if a.element == b.element || b.element == c.element || a.element == c.element {
					searchXTarget(filter(a, b, c), a, b, c)
				}
			}

			if p.Y < len(chessboard)-2 {
				a, b, c := p, point{p.X, p.Y + 1, chessboard[p.X][p.Y+1]}, point{p.X, p.Y + 2, chessboard[p.X][p.Y+2]}
				if a.element == b.element || b.element == c.element || a.element == c.element {
					searchYTarget(filter(a, b, c), a, b, c)
				}
			}
		}
	}
	fmt.Println("targetPoint", targetPoint)
}

func chessDrop() {
	// TODO: 下落队列

	chessboard[6][6] = "X"

	// dropQueue := []point

	for row := len(chessboard) - 2; row >= 0; row-- {
		for col := len(chessboard[row]) - 1; col >= 0; col-- {

			if chessboard[row+1][col] == "X" {
				chessboard[row+1][col] = chessboard[row][col]
				chessboard[row][col] = "X"
			}

			//	append(dropQueue, point{})

			//fmt.Println("chessboard[col]", chessboard[row][col])
			// chessboard[0][0] = chessboard[1][0]
		}
	}

	fmt.Println()
	for c := range chessboard {
		fmt.Println(chessboard[c])
	}
}

func checkPos() {
loop:
	for row, cb := range chessboard {
		for col := range cb {
			p := point{row, col, chessboard[row][col]}
			if p.X < len(chessboard)-2 {
				a, b, c := chessboard[p.X+1][p.Y], chessboard[p.X+2][p.Y], chessboard[p.X][p.Y]
				if a == b && b == c {
					fmt.Println("消除 X", a, b, c)
					chessboard[p.X+1][p.Y], chessboard[p.X+2][p.Y], chessboard[p.X][p.Y] = "X", "X", "X"

					for i := p.X + 3; i < len(chessboard); i++ {
						if chessboard[i][p.Y] == a {
							fmt.Println("继续消除 X", chessboard[i][p.Y], i, p.Y)
							chessboard[i][p.Y] = "X"
						} else {
							break
						}
					}

					initChessboard(chessboard)

					goto loop
				}
			}

			if p.Y < len(chessboard[0])-2 {
				a, b, c := chessboard[p.X][p.Y+1], chessboard[p.X][p.Y+2], chessboard[p.X][p.Y]
				if a == b && b == c {
					fmt.Println("消除 Y", a, b, c)
					chessboard[p.X][p.Y+1], chessboard[p.X][p.Y+2], chessboard[p.X][p.Y] = "X", "X", "X"

					for i := p.Y + 3; i < len(chessboard[0]); i++ {
						if chessboard[p.X][i] == a {
							fmt.Println("继续消除 Y", chessboard[p.X][i], p.X, i)
							chessboard[p.X][i] = "X"
						} else {
							break
						}
					}

					initChessboard(chessboard)

					goto loop
				}
			}
		}
	}
}

func searchXTarget(f, a, b, c point) {
	switch f.element {
	case a.element:
		if f.X != 0 {
			if chessboard[f.X-1][f.Y] == b.element {
				targetPoint[f] = make(map[point][]point)
				targetPoint[f][point{f.X - 1, f.Y, chessboard[f.X-1][f.Y]}] = []point{b, c}
			}
		}

		if f.Y != len(chessboard)-1 {
			if chessboard[f.X][f.Y+1] == b.element {
				targetPoint[f] = make(map[point][]point)
				targetPoint[f][point{f.X, f.Y + 1, chessboard[f.X][f.Y+1]}] = []point{b, c}
			}
		}

		if f.Y != 0 {
			if chessboard[f.X][f.Y-1] == b.element {
				targetPoint[f] = make(map[point][]point)
				targetPoint[f][point{f.X, f.Y - 1, chessboard[f.X][f.Y-1]}] = []point{b, c}
			}
		}
	case b.element:
		if f.Y != len(chessboard)-1 {
			if chessboard[f.X][f.Y+1] == a.element {
				targetPoint[f] = make(map[point][]point)
				targetPoint[f][point{f.X, f.Y + 1, chessboard[f.X][f.Y+1]}] = []point{a, c}
			}
		}

		if f.Y != 0 {
			if chessboard[f.X][f.Y-1] == a.element {
				targetPoint[f] = make(map[point][]point)
				targetPoint[f][point{f.X, f.Y - 1, chessboard[f.X][f.Y-1]}] = []point{a, c}
			}
		}
	case c.element:
		if f.X != len(chessboard)-1 {
			if chessboard[f.X+1][f.Y] == a.element {
				targetPoint[f] = make(map[point][]point)
				targetPoint[f][point{f.X + 1, f.Y, chessboard[f.X+1][f.Y]}] = []point{a, b}
			}
		}

		if f.Y != len(chessboard)-1 {
			if chessboard[f.X][f.Y+1] == a.element {
				targetPoint[f] = make(map[point][]point)
				targetPoint[f][point{f.X, f.Y + 1, chessboard[f.X][f.Y+1]}] = []point{a, b}
			}
		}

		if f.Y != 0 {
			if chessboard[f.X][f.Y-1] == a.element {
				targetPoint[f] = make(map[point][]point)
				targetPoint[f][point{f.X, f.Y - 1, chessboard[f.X][f.Y-1]}] = []point{a, b}
			}
		}
	}
}

func searchYTarget(f, a, b, c point) {
	switch f.element {
	case a.element:
		if f.Y != 0 {
			if chessboard[f.X][f.Y-1] == b.element {
				targetPoint[f] = make(map[point][]point)
				targetPoint[f][point{f.X, f.Y - 1, chessboard[f.X][f.Y-1]}] = []point{b, c}
			}
		}

		if f.X != 0 {
			if chessboard[f.X-1][f.Y] == b.element {
				targetPoint[f] = make(map[point][]point)
				targetPoint[f][point{f.X - 1, f.Y, chessboard[f.X-1][f.Y]}] = []point{b, c}
			}
		}

		if f.X != len(chessboard)-1 {
			if chessboard[f.X+1][f.Y] == b.element {
				targetPoint[f] = make(map[point][]point)
				targetPoint[f][point{f.X + 1, f.Y, chessboard[f.X+1][f.Y]}] = []point{b, c}
			}
		}
	case b.element:
		if f.X != 0 {
			if chessboard[f.X-1][f.Y] == a.element {
				targetPoint[f] = make(map[point][]point)
				targetPoint[f][point{f.X - 1, f.Y, chessboard[f.X-1][f.Y]}] = []point{a, c}
			}
		}

		if f.X != len(chessboard)-1 {
			if chessboard[f.X+1][f.Y] == a.element {
				targetPoint[f] = make(map[point][]point)
				targetPoint[f][point{f.X + 1, f.Y, chessboard[f.X+1][f.Y]}] = []point{a, c}
			}
		}
	case c.element:
		if f.Y != len(chessboard)-1 {
			if chessboard[f.X][f.Y+1] == b.element {
				targetPoint[f] = make(map[point][]point)
				targetPoint[f][point{f.X, f.Y + 1, chessboard[f.X][f.Y+1]}] = []point{a, b}
			}
		}

		if f.X != 0 {
			if chessboard[f.X-1][f.Y] == b.element {
				targetPoint[f] = make(map[point][]point)
				targetPoint[f][point{f.X - 1, f.Y, chessboard[f.X-1][f.Y]}] = []point{a, b}
			}
		}

		if f.X != len(chessboard)-1 {
			if chessboard[f.X+1][f.Y] == b.element {
				targetPoint[f] = make(map[point][]point)
				targetPoint[f][point{f.X + 1, f.Y, chessboard[f.X+1][f.Y]}] = []point{a, b}
			}
		}
	}
}

func initChessboard(chessboard [][]string) {
	for _, chess := range chessboard {
		for index, c := range chess {
			if c == "X" {
				i := rand.Intn(5)
				chess[index] = element[i]
			}
		}
	}
}
