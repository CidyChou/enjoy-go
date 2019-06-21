package main

import "fmt"

type Game struct{}

type IGame interface {
	Name() string
}

func (g *Game) play(game IGame) {
	fmt.Println(fmt.Sprintf("%s is awesome!", game.Name()))
}

type LOL struct {
	Game
}

type Dota struct {
	Game
}

func (d *Dota) Name() string {
	return "Dota"
}

func (l *LOL) Name() string {
	return "LOL"
}

func main() {
	dota := &Dota{}
	dota.play(dota)
	lol := &LOL{}
	lol.play(lol)

	fmt.Println(dota.Name())
}
