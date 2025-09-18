package main

import (
	"bytes"
	"fmt"
	"os"
)

const (
	NOTHING = 0
	WALL    = 1
	PLAYER  = 69
)

type Game struct {
	isRunning bool
	level     *level
	drawbuf   *bytes.Buffer
}
type level struct {
	height, width int
	data          [][]byte
}

func NewLevel(width, height int) *level {
	data := make([][]byte, height)
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			data[h] = make([]byte, width)

		}
	}
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if h == 0 {
				data[h][w] = WALL
			}
			if w == 0 {
				data[h][w] = WALL
			}
			if h == height-1 {
				data[h][w] = WALL
			}
			if w == width-1 {
				data[h][w] = WALL
			}

		}
	}
	return &level{
		data: data,
	}
}
func NewGame(width, height int) *Game {
	lvl := NewLevel(width, height)
	return &Game{
		level:   lvl,
		drawbuf: new(bytes.Buffer),
	}
}
func (g *Game) Start() {
	g.isRunning = true
	g.loop()
}
func (g *Game) loop() {
	for g.isRunning {
		g.update()
		g.render()
	}
}
func (g *Game) renderarena() {
	for h := 0; h < g.level.height; h++ {
		for w := 0; w < g.level.width; w++ {
			if g.level.data[h][w] == NOTHING {
				g.drawbuf.WriteString(" ")
			}
			if g.level.data[h][w] == WALL {
				g.drawbuf.WriteString("#")
			}

		}
		g.drawbuf.WriteString("\n")
	}
}
func (g *Game) update() {}
func (g *Game) render() {
	g.renderarena()
	fmt.Fprint(os.Stdout, g.drawbuf.String())
}

func main() {
	height := 15
	width := 40
	fmt.Println("main")
	g := NewGame(width, height)
	g.Start()

}
