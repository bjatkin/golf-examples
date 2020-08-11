package main

import (
	"fantasyConsole/golf"
	"fmt"
)

var playerXY = vec2{200, 132}
var playerSpeed = 2.0
var playerFacing int
var playerWalking bool
var playerOpt = golf.SOp{TCol: golf.Col6, W: 2, H: 2}

const (
	facingDown = iota
	facingUp
	facingLeft
	facingRight
)

func updatePlayer() {
	playerWalking = false
	if g.Btnp(golf.TKey) {
		fmt.Printf("%v\n", playerXY)
	}
	if g.Btn(golf.WKey) {
		playerFacing = facingUp
		playerXY.y -= playerSpeed
		if mapCollide(playerXY) {
			playerXY.y += playerSpeed
		}
		playerWalking = true
		playerOpt.FH = false
	}
	if g.Btn(golf.SKey) {
		playerFacing = facingDown
		playerXY.y += playerSpeed
		if mapCollide(playerXY) {
			playerXY.y -= playerSpeed
		}
		playerWalking = true
		playerOpt.FH = false
	}
	if g.Btn(golf.AKey) {
		playerFacing = facingLeft
		playerXY.x -= playerSpeed
		if mapCollide(playerXY) {
			playerXY.x += playerSpeed
		}
		playerWalking = true
		playerOpt.FH = false
	}
	if g.Btn(golf.DKey) {
		playerFacing = facingRight
		playerXY.x += playerSpeed
		if mapCollide(playerXY) {
			playerXY.x -= playerSpeed
		}
		playerWalking = true
		playerOpt.FH = true
	}
}

func drawPlayer() {
	index := 2
	if playerFacing == facingUp {
		index = 8
	}
	if playerFacing == facingLeft || playerFacing == facingRight {
		index = 14
	}
	if playerWalking {
		index += 2 * (((g.Frames() / 5) % 3) - 1)
	}
	g.Spr(index, playerXY.x, playerXY.y, playerOpt)
}

func mapCollide(p vec2) bool {
	if p.x < -1 || p.y < -1 || p.x > mainScean.mapWH.x*8-14 || p.y > mainScean.mapWH.y*8-16 {
		return true
	}
	// tile offsets
	x := int(mainScean.mapXY.x)
	y := int(mainScean.mapXY.y)

	// 4 corners
	a := g.Fget(g.Mget(int((p.x+2)/8)+x, int((p.y+2)/8))+y, 0)
	b := g.Fget(g.Mget(int((p.x+14)/8)+x, int((p.y+2)/8))+y, 0)
	c := g.Fget(g.Mget(int((p.x+2)/8)+x, int((p.y+16)/8))+y, 0)
	d := g.Fget(g.Mget(int((p.x+14)/8)+x, int((p.y+16)/8))+y, 0)

	// vertical center
	e := g.Fget(g.Mget(int((p.x+2)/8)+x, int((p.y+8)/8))+y, 0)
	f := g.Fget(g.Mget(int((p.x+14)/8)+x, int((p.y+8)/8))+y, 0)

	// horizontal center
	h := g.Fget(g.Mget(int((p.x+8)/8)+x, int((p.y+2)/8))+y, 0)
	i := g.Fget(g.Mget(int((p.x+8)/8)+x, int((p.y+16)/8))+y, 0)

	return (a || b || c || d || e || f || h || i)
}
