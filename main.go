package main

import "fmt"

const (
	NOTHING = 0
	WALL    = 1
	PLAYER  = 69
)

type Game struct{}

func (g *Game) update() {}
func (g *Game) render() {

}

func main() {
	height := 5
	width := 5
	level := make([][]byte, height)
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			level[h] = make([]byte, width)

		}
	}
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			level[h][w] = WALL

		}
	}

	fmt.Println("level:", level)
}
