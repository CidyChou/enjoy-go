package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

type point struct {
	X int
	Y int
	//element string
}

var chessboard [][]string
var element []string
var score int

func init() {
	chessboard = [][]string{
		{"Q", "Q", "Q", "Q", "X", "X", "X", "X"},
		{"Q", "X", "X", "X", "X", "X", "X", "X"},
		{"Q", "X", "X", "X", "X", "X", "X", "X"},
		{"Q", "X", "X", "X", "X", "X", "X", "X"},
		{"X", "X", "X", "X", "X", "X", "X", "X"},
		{"X", "X", "X", "X", "X", "X", "X", "X"},
		{"X", "X", "X", "X", "X", "X", "X", "X"},
		{"X", "X", "X", "X", "X", "X", "X", "X"},
	}

	//element = []string{"☄", "✹", "✪", "✢", "✭"}
	element = []string{"H", "%", "&", "Q", "M"}

	initChessboard(chessboard)

	fmt.Println("--------------------游戏开始-------------------")
	fmt.Println()
	fmt.Println()
	for c := range chessboard {
		fmt.Println("         ", chessboard[c])
	}
	fmt.Println()
}

func main() {
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
			mdScore := score
			chessboard[int(X1)][int(Y1)], chessboard[int(X2)][int(Y2)] = chessboard[int(X2)][int(Y2)], chessboard[int(X1)][int(Y1)]

			onTrigger()

			if mdScore == score {
				fmt.Println("无消除")
				chessboard[int(X2)][int(Y2)], chessboard[int(X1)][int(Y1)] = chessboard[int(X1)][int(Y1)], chessboard[int(X2)][int(Y2)]
			}

			// fmt.Println("--------------------移动后-------------------")
			// fmt.Println()

			// for c := range chessboard {
			// 	fmt.Println("         ", chessboard[c])
			// }
			// fmt.Println()
		}
	}
}

func onTrigger() {

	ScanPos()

	fmt.Println()
	for c := range chessboard {
		fmt.Println("         ", chessboard[c])
	}
	fmt.Println()

}

// ScanPos is 扫描棋盘
func ScanPos() bool {
	bol := true

	for bol {
		bol = false
		for row, cb := range chessboard {
			for col := range cb {
				p := point{row, col}
				if p.X < len(chessboard)-2 {
					a, b, c := chessboard[p.X+1][p.Y], chessboard[p.X+2][p.Y], chessboard[p.X][p.Y]
					if a == b && b == c {
						bol = true
						score += 100
						fmt.Println("消除", a, b, c, "积分+100;Score:", score)

						chessboard[p.X+1][p.Y], chessboard[p.X+2][p.Y], chessboard[p.X][p.Y] = "X", "X", "X"

						for i := p.X + 2; i < len(chessboard); i++ {
							if chessboard[i][p.Y] == a {
								fmt.Println("消除", chessboard[i][p.Y], i, p.Y)
								chessboard[i][p.Y] = "X"
							}
						}

					}
				}

				if p.Y < len(chessboard[0])-2 {
					a, b, c := chessboard[p.X][p.Y+1], chessboard[p.X][p.Y+2], chessboard[p.X][p.Y]
					if a == b && b == c {
						bol = true
						score += 100
						fmt.Println("消除", a, b, c, "积分+100;Score:", score)
						chessboard[p.X][p.Y+1], chessboard[p.X][p.Y+2], chessboard[p.X][p.Y] = "X", "X", "X"

						for i := p.Y + 2; i < len(chessboard[0]); i++ {
							if chessboard[p.X][i] == a {
								fmt.Println("消除", chessboard[p.X][i], p.X, i)
								chessboard[p.X][i] = "X"
							}
						}

					}
				}

				// for c := range chessboard {
				// 	fmt.Println("         ", chessboard[c])
				// }
				// fmt.Println()
			}
		}
		initChessboard(chessboard)
	}

	return bol
}

func initChessboard(chessboard [][]string) {
	for _, chess := range chessboard {
		for index, c := range chess {
			if c == "X" {
				//rand.Seed(time.Now().UnixNano())
				i := rand.Intn(5)
				// i, _ := rand.Int(rand.Reader, big.NewInt(100))
				// fmt.Println("i", i)
				// var ii int
				// ii = *i
				chess[index] = element[i]
			}
		}
	}
}
