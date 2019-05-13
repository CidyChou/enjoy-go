package main

import (
	"fmt"
	"math/rand"
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
	for row, cb := range chessboard {
		for col := range cb {
			p := point{row, col}
			ScanPos(&p)
		}
	}

	for c := range chessboard {
		fmt.Println(chessboard[c])
	}
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

// ScanPos is 扫描棋盘
func ScanPos(p *point) {
	if p.X < len(chessboard)-2 {
		a, b, c := chessboard[p.X+1][p.Y], chessboard[p.X+2][p.Y], chessboard[p.X][p.Y]
		if a == b && b == c {
			fmt.Println("X消除 ------")
			chessboard[p.X+1][p.Y], chessboard[p.X+2][p.Y], chessboard[p.X][p.Y] = "X", "X", "X"
		}
	}

	if p.Y < len(chessboard[0])-2 {
		a, b, c := chessboard[p.X][p.Y+1], chessboard[p.X][p.Y+2], chessboard[p.X][p.Y]
		if a == b && b == c {
			fmt.Println("Y消除 ------")
			chessboard[p.X][p.Y+1], chessboard[p.X][p.Y+2], chessboard[p.X][p.Y] = "X", "X", "X"
		}
	}
}
