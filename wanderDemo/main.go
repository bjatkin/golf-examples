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

	initEventHandler()
	initSceans()
	initCharacterEvents()
}

func update() {
	g.PalA(3)
	g.PalB(10)

	updateCamera()
	updatePlayer()

	newScean, exited := exitScean(playerXY)
	if exited {
		prevScean = mainScean
		mainScean = newScean
		transitionStart = g.Frames()
		g.Update = updateTransition
		g.Draw = drawTransition
	}

	if g.Btnp(golf.ZKey) {
		poi, interact := interactWithScean(playerXY)
		if interact {
			poi.resetInteraction()
			mainInteraction = poi
			g.Update = updateInteraction
			g.Draw = drawInteraction
		}
	}

	if g.Frames() == 253*30 {
		storyEventHandler.sendEvent(twoFiftySixSecondsPassed)
	}
}

func updateCamera() {
	cx, cy := playerXY.x-88, playerXY.y-88
	if cx < 0 && mainScean.mapWH.x >= 16 {
		cx = 0
	}
	if cy < 0 && mainScean.mapWH.y >= 16 {
		cy = 0
	}
	if cx > (mainScean.mapWH.x-24)*8 && mainScean.mapWH.x >= 16 {
		cx = (mainScean.mapWH.x - 24) * 8
	}
	if cy > (mainScean.mapWH.y-24)*8 && mainScean.mapWH.y >= 16 {
		cy = (mainScean.mapWH.y - 24) * 8
	}
	g.Camera(int(cx), int(cy))
}

func draw() {
	g.Cls(golf.Col4)
	mainScean.draw()
	drawPlayer()
	mainScean.drawPOI(playerXY)
}

func updateTransition() {
	updateCamera()
}

var transitionStart = 0

func drawTransition() {
	f := float64(g.Frames()-transitionStart) * 2
	if f <= 96 {
		g.RectFill(f-96, 0, 96, 192, golf.Col0, true)
		g.RectFill(192-f, 0, 96, 192, golf.Col0, true)
	}
	if f > 96 {
		g.Cls(golf.Col4)
		mainScean.draw()
		playerXY = mainScean.entrances[prevScean][1]
		playerWalking = false
		drawPlayer()

		g.RectFill(-(f - 96), 0, 96, 192, golf.Col0, true)
		g.RectFill(96+(f-96), 0, 96, 192, golf.Col0, true)
	}
	if f > 192 {
		g.Update = update
		g.Draw = draw
	}
}
