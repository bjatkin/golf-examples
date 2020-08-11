package main

import "fantasyConsole/golf"

var playerXY = vec2{}
var playerFacing int
var playerWalking bool
var playerOpt = golf.SOp{TCol: golf.Col5, W: 2, H: 2}

const (
	facingDown = iota
	facingUp
	facingLeft
	facingRight
)

func updatePlayer() {
	playerWalking = false
	if g.Btnp(golf.WKey) {
		playerFacing = facingUp
		playerXY.y--
		playerWalking = true
	}
	if g.Btnp(golf.SKey) {
		playerFacing = facingDown
		playerXY.y++
		playerWalking = true
	}
	if g.Btnp(golf.AKey) {
		playerFacing = facingLeft
		playerXY.x--
		playerWalking = true
	}
	if g.Btnp(golf.DKey) {
		playerFacing = facingRight
		playerXY.x++
		playerWalking = true
	}
}

func drawPlayer() {
	index := 0
	if playerFacing == facingUp {
		index = 2
	}
	if playerFacing == facingLeft {
		index = 4
	}
	if playerFacing == facingRight {
		index = 5
	}
	if playerWalking {
		index += 2 * (((g.Frames() / 30) % 3) - 1)
	}
	g.Spr(index, playerXY.x, playerXY.y, playerOpt)
}
