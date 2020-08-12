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

func initGopher() {
	for x := 0; x < 128; x++ {
		for y := 0; y < 128; y++ {
			if g.Mget(x, y) == 201 { // The gopher sprite

				// erase the gopher tiles
				g.Mset(x, y, 0)
				g.Mset(x+1, y, 0)
				g.Mset(x, y+1, 0)
				g.Mset(x+1, y+1, 0)

				// replace it the gopher character
				newMob(float64(x*8), float64(y*8), []drawable{
					&ani{ //Waiting
						frames: []int{201, 203},
						speed:  0.05,
						o:      golf.SOp{W: 2, H: 2},
					},
					&ani{ //Talking
						frames: []int{137, 201},
						speed:  0.05,
						o:      golf.SOp{W: 2, H: 2},
					},
				})
				break
			}
		}
	}
}
