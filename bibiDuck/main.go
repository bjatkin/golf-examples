package main

import (
	"fantasyConsole/golf"
)

var g *golf.Engine
var alpha = golf.SOp{TCol: golf.Col5}
var player multiAniSprite

func main() {
	g = golf.NewEngine(update, draw)

	g.LoadMap(mapData)
	g.LoadSprs(spriteSheet)
	g.PalA(golf.Pal7)
	g.PalB(golf.Pal8)

	g.RAM[0x3601] = 0
	initPickups()
	player = multiAniSprite{
		id: 0,
		states: []aniSprite{
			aniSprite{
				id:     0,
				frames: []int{1, 3},
				speed:  0.03,
				x:      40,
				y:      161,
				o:      golf.SOp{TCol: golf.Col5, W: 2, H: 2},
			},
		},
	}
	allAni = append(allAni, &player.states[0])

	g.Run()
}

var drawMap bool
var cameraX, cameraY int

func update() {
	g.Camera(cameraX, cameraY)
}

func draw() {
	g.Cls(golf.Col5)
	g.Map(0, 0, 128, 128, 0, 0, alpha)

	drawAniSprites()
}

func initPickups() {
	for x := 0; x < 128; x++ {
		for y := 0; y < 128; y++ {
			tile := g.Mget(x, y)
			if tile == 32 {
				g.Mset(x, y, 0)
				newFether(float64(x*8), float64(y*8))
			}
			if tile == 193 {
				g.Mset(x, y, 0)
				g.Mset(x+1, y, 0)
				g.Mset(x, y+1, 0)
				g.Mset(x+1, y+1, 0)
				newEgg(float64(x*8), float64(y*8))
			}
		}
	}
}
