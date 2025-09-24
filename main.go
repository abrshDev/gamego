package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	NOTHING     = 0
	WALL        = 1
	PLAYER      = 69
	MAX_SAMPLES = 100
)

type position struct {
	x, y int
}
type player struct {
	level   *level
	pos     position
	reverse bool
	input   *input
}
type Stats struct {
	start  time.Time
	frames int
	fps    float64
}

type Game struct {
	isRunning bool
	level     *level
	player    *player
	drawbuf   *bytes.Buffer
	stats     *Stats
	input     *input
}
type level struct {
	height, width int
	data          [][]int
}
type input struct {
	pressedkey byte
}

func (i *input) update() {
	/* i.pressedkey = 0

			tick := time.NewTicker(time.Millisecond * 2)

	free: */

	/* for {
		select {
		case <-tick.C:
			break free
		default:
			b := make([]byte, 1)
			os.Stdin.Read(b)
			i.pressedkey = b[0]
		}
	} */

}

func (p *player) update() {
	if p.reverse {
		p.pos.x -= 1

		if p.pos.x == 2 {
			p.pos.x += 1
			p.reverse = false
		}
		return
	}
	p.pos.x += 1
	if p.pos.x == p.level.width-2 {
		p.pos.x -= 1
		p.reverse = true
	}
}
func NewStats() *Stats {
	return &Stats{
		start: time.Now(),
		fps:   62,
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
	data := make([][]int, height)
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			data[h] = make([]int, width)

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
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo", "min", "1").Run()
	lvl := NewLevel(width, height)
	inp := &input{}
	return &Game{
		level:   lvl,
		drawbuf: new(bytes.Buffer),
		input:   inp,
		player: &player{
			level: lvl,
			pos: position{
				x: 8, y: 5,
			},
			input: inp,
		},
		stats: NewStats(),
	}
}
func (g *Game) Start() {
	g.isRunning = true
	g.loop()
}
func (g *Game) loop() {
	for g.isRunning {
		g.input.update()
		g.update()

		g.render()
		g.stats.update()
		time.Sleep(time.Millisecond * 16) //limit fps

	}
}

func (l *level) set(pos position, v int) {
	l.data[pos.y][pos.x] = v
}
func (g *Game) renderarena() {
	for h := 0; h < g.level.height; h++ {
		for w := 0; w < g.level.width; w++ {
			if g.level.data[h][w] == NOTHING {
				g.drawbuf.WriteString(" ")
			}
			if g.level.data[h][w] == WALL {
				g.drawbuf.WriteString("✦")
			}
			if g.level.data[h][w] == PLAYER {
				g.drawbuf.WriteString("🎮")
			}

		}
		g.drawbuf.WriteString("\n")
	}
}
func (g *Game) renderstats() {
	g.drawbuf.WriteString(" --STATS--\n")
	g.drawbuf.WriteString(fmt.Sprintf("fps:%.2f\n", g.stats.fps))
	g.drawbuf.WriteString(fmt.Sprintf("pressedkey:%v\n", g.input.pressedkey))
}
func (g *Game) update() {
	g.level.set(g.player.pos, NOTHING)
	g.player.update()
	g.level.set(g.player.pos, PLAYER)
}
func (g *Game) render() {
	g.drawbuf.Reset()
	fmt.Fprint(os.Stdout, "\033[2J\033[1;1H")
	g.renderarena()
	g.renderstats()
	fmt.Fprint(os.Stdout, g.drawbuf.String())

}

func main() {
	height := 15
	width := 80
	g := NewGame(width, height)
	g.Start()

}
