package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

const (
	NOTHING     = 0
	WALL        = 1
	PLAYER      = 69
	MAX_SAMPLES = 100
)

type Stats struct {
	start  time.Time
	frames int
	fps    float64
}

type Game struct {
	isRunning bool
	level     *level
	drawbuf   *bytes.Buffer
	stats     *Stats
}
type level struct {
	height, width int
	data          [][]byte
}

func NewStats() *Stats {
	return &Stats{
		start: time.Now(),
	}
}
func (s *Stats) update() {
	s.frames++
	if s.frames == MAX_SAMPLES {
		s.fps = float64(s.frames) / time.Since(s.start).Seconds()
		s.frames = 0
		s.start = time.Now()
	}
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
		width:  width,
		height: height,
		data:   data,
	}
}
func NewGame(width, height int) *Game {
	lvl := NewLevel(width, height)
	return &Game{
		level:   lvl,
		drawbuf: new(bytes.Buffer),
		stats:   NewStats(),
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
		g.stats.update()
		time.Sleep(time.Millisecond * 14)

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
func (g *Game) renderstats() {
	g.drawbuf.WriteString(" --STATS--\n")
	g.drawbuf.WriteString(fmt.Sprintf("fps:%.2f", g.stats.fps))
}
func (g *Game) update() {}
func (g *Game) render() {
	g.drawbuf.Reset()
	fmt.Fprint(os.Stdout, "\033[2J\033[1;1H")
	g.renderarena()
	g.renderstats()
	fmt.Fprint(os.Stdout, g.drawbuf.String())

}

func main() {
	height := 15
	width := 40
	fmt.Println("main")
	g := NewGame(width, height)
	g.Start()

}
