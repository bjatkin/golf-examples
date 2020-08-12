package main

import (
	"fantasyConsole/golf"
)

var g *golf.Engine
var alpha = golf.SOp{TCol: golf.Col5}

func main() {
	g = golf.NewEngine(update, draw)

	g.LoadMap(mapData)
	g.LoadSprs(spriteSheet)
	g.LoadFlags(spriteFlags)

	g.PalA(golf.Pal7)
	g.PalB(golf.Pal8)

	initGame()

	g.Run()
}

func initGame() {
	initPickups()
	initPlayer()
	initGopher()
	initHUD()
	initConvo()
	initParticles()
}

var cameraX, cameraY int

func update() {
	// The player should only be able to move
	// if they are not in a conversation
	if !gopherConvo.running {
		updatePlayer()
	}

	updateCamera()
	checkPickups()

	if g.Btnp(golf.XKey) && nearGopher {
		// Make sure that the duck faces the gopher
		if duck.x < 387 {
			duck.a.(*stateAni).state = waitRight
		} else {
			duck.a.(*stateAni).state = waitLeft
		}
		// Start/ move to the next line in the conversation
		gopherConvo.next()
	}
}

func updateCamera() {
	// move the camera with the player
	// but keey it in bounds

	cameraX = int(duck.x) - 96
	if cameraX < 0 {
		cameraX = 0
	}
	if cameraX > 422 {
		cameraX = 422
	}
	g.Camera(cameraX, cameraY)
}

func draw() {
	g.Cls(golf.Col5)
	g.Map(0, 0, 128, 32, 0, 0, alpha)

	if !gopherConvo.running {
		// Prevent the player state from updating
		// while talking to joe gopher
		drawPlayer()
	}

	drawSprites()
	drawHUD()

	gopherConvo.draw()

	drawSpeachBubble()
	drawParticles()
}
