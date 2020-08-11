package main

import (
	"fantasyConsole/golf"
)

var g *golf.Engine

func main() {
	g = golf.NewEngine(update, draw)

	g.LoadSprs(spriteSheet)
	g.LoadMap(mapData)
	g.LoadFlags(spriteFlags)
	g.PalA(2)
	g.PalB(3)

	g.Run()
}

var cameraX, cameraY = 0, 0

func update() {
	g.PalA(3)
	g.PalB(10)
	if g.Btn(golf.WKey) {
		cameraY--
	}
	if g.Btn(golf.SKey) {
		cameraY++
	}
	if g.Btn(golf.DKey) {
		cameraX++
	}
	if g.Btn(golf.AKey) {
		cameraX--
	}
	if cameraX < 0 {
		cameraX = 0
	}
	if cameraX > 832 {
		cameraX = 832
	}
	if cameraY < 0 {
		cameraY = 0
	}
	if cameraY > 832 {
		cameraY = 832
	}
	g.Camera(cameraX, cameraY)

	newScean, exited := exitScean(playerXY)
	if exited {
		playerXY = newScean.entrances[mainScean]
		mainScean = newScean
	}
}

func draw() {
	g.Cls(golf.Col7)
	mainScean.draw()
}
