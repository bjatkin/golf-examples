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
	initGame()

	g.Run()
}

func initGame() {
	// Pallet for the starting logo
	g.PalA(2)
	g.PalB(3)

	initSceans()
}

func update() {
	g.PalA(3)
	g.PalB(10)
	cx, cy := int(playerXY.x)-88, int(playerXY.y)-88
	if cx < 0 && mainScean.mapWH.x >= 16 {
		cx = 0
	}
	if cy < 0 && mainScean.mapWH.y >= 16 {
		cy = 0
	}
	if cx > int((mainScean.mapWH.x-24)*8) && mainScean.mapWH.x >= 16 {
		cx = int((mainScean.mapWH.x - 24) * 8)
	}
	if cy > int((mainScean.mapWH.y-24)*8) && mainScean.mapWH.y >= 16 {
		cy = int((mainScean.mapWH.y - 24) * 8)
	}
	g.Camera(cx, cy)

	updatePlayer()

	newScean, exited := exitScean(playerXY)
	if exited {
		playerXY = newScean.entrances[mainScean][1]
		mainScean = newScean
	}
}

func draw() {
	g.Cls(golf.Col4)
	mainScean.draw()
	drawPlayer()
	if g.Frames()%2 == 0 && false {
		for x := 0.0; x < mainScean.mapWH.x; x++ {
			for y := 0.0; y < mainScean.mapWH.y; y++ {
				if g.Fget(g.Mget(int(x+mainScean.mapXY.x), int(y+mainScean.mapXY.y)), 0) {
					g.RectFill(x*8, y*8, 8, 8, golf.Col0)
				}
			}
		}
	}
}
