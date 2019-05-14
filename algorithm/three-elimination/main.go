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

	element = []string{"✭", "☄", "✹", "✪", "✢"}

	initChessboard(chessboard)

	for c := range chessboard {
		fmt.Println(chessboard[c])
	}
}

func main() {
	for {

		onTrigger()

		var p1, p2 string
		fmt.Println("请输入要移动的棋子的位置 与 移动后的位置:")
		fmt.Scanln(&p1, &p2)

		X1, _ := strconv.ParseFloat(strings.Split(p1, ",")[0], 64)
		Y1, _ := strconv.ParseFloat(strings.Split(p1, ",")[1], 64)

		X2, _ := strconv.ParseFloat(strings.Split(p2, ",")[0], 64)
		Y2, _ := strconv.ParseFloat(strings.Split(p2, ",")[1], 64)

		if math.Abs(X2-X1) > 1 || math.Abs(Y2-Y1) > 1 {
			fmt.Println("请输入正确的位置")
		} else {
			chessboard[int(X1)][int(Y1)], chessboard[int(X2)][int(Y2)] = chessboard[int(X2)][int(Y2)], chessboard[int(X1)][int(Y1)]
			for c := range chessboard {
				fmt.Println(chessboard[c])
			}
		}
	}
}

func onTrigger() {

	for row, cb := range chessboard {
		for col := range cb {
			p := point{row, col}
			ScanPos(&p)
		}
	}

	initChessboard(chessboard)

	for c := range chessboard {
		fmt.Println(chessboard[c])
	}

}

// ScanPos is 扫描棋盘
func ScanPos(p *point) bool {
	var bol bool
	if p.X < len(chessboard)-2 {
		a, b, c := chessboard[p.X+1][p.Y], chessboard[p.X+2][p.Y], chessboard[p.X][p.Y]
		if a == b && b == c {
			bol = true
			fmt.Println("X消除 ------")
			chessboard[p.X+1][p.Y], chessboard[p.X+2][p.Y], chessboard[p.X][p.Y] = "X", "X", "X"
		}
	}

	if p.Y < len(chessboard[0])-2 {
		a, b, c := chessboard[p.X][p.Y+1], chessboard[p.X][p.Y+2], chessboard[p.X][p.Y]
		if a == b && b == c {
			bol = true
			fmt.Println("Y消除 ------")
			chessboard[p.X][p.Y+1], chessboard[p.X][p.Y+2], chessboard[p.X][p.Y] = "X", "X", "X"
		}
	}
	return bol
}

func initChessboard(chessboard [][]string) {
	for _, chess := range chessboard {
		for index, c := range chess {
			if c == "X" {
				//rand.Seed(time.Now().UnixNano())
				i := rand.Intn(4)
				// i, _ := rand.Int(rand.Reader, big.NewInt(100))
				// fmt.Println("i", i)
				// var ii int
				// ii = *i
				chess[index] = element[i]
			}
		}
	}
}
