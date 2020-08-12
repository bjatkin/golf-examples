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

	dx, dy := 0.0, 0.0
	if g.Btn(golf.UpArrow) {
		playerFacing = facingUp
		dy = -playerSpeed
		playerWalking = true
		playerOpt.FH = false
	}
	if g.Btn(golf.DownArrow) {
		playerFacing = facingDown
		dy = playerSpeed
		playerWalking = true
		playerOpt.FH = false
	}
	if g.Btn(golf.LeftArrow) && dy == 0 {
		playerFacing = facingLeft
		dx = -playerSpeed
		playerWalking = true
		playerOpt.FH = false
	}
	if g.Btn(golf.RightArrow) && dy == 0 {
		playerFacing = facingRight
		dx = playerSpeed
		playerWalking = true
		playerOpt.FH = true
	}

	movePlayer(dx, dy)
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

func movePlayer(dx, dy float64) {
	dest := vec2{playerXY.x + dx, playerXY.y + dy}
	mapCollide, shiftX, shiftY := mapCollide(dest, playerFacing)
	sceanCollide := sceanCollide(dest)
	if !mapCollide && !sceanCollide {
		playerXY.x += dx
		playerXY.y += dy
	}
	if !sceanCollide && mapCollide && shiftX != 0 {
		playerXY.x += shiftX
	}
	if !sceanCollide && mapCollide && shiftY != 0 {
		playerXY.y += shiftY
	}
}

func mapCollide(p vec2, facing int) (bool, float64, float64) {
	if p.x < -1 || p.y < -1 || p.x > mainScean.mapWH.x*8-14 || p.y > mainScean.mapWH.y*8-16 {
		return true, 0, 0
	}
	// tile offsets
	x := int(mainScean.mapXY.x)
	y := int(mainScean.mapXY.y)

	if facing == facingUp {
		a := g.Fget(g.Mget(int((p.x+2)/8)+x, int((p.y+2)/8))+y, 0)
		b := g.Fget(g.Mget(int((p.x+8)/8)+x, int((p.y+2)/8))+y, 0)
		c := g.Fget(g.Mget(int((p.x+14)/8)+x, int((p.y+2)/8))+y, 0)
		shift := 0.0
		if !a && c {
			shift = -1
		}
		if !c && a {
			shift = 1
		}
		return (a || b || c), shift, 0
	}

	if facing == facingDown {
		a := g.Fget(g.Mget(int((p.x+2)/8)+x, int((p.y+16)/8))+y, 0)
		b := g.Fget(g.Mget(int((p.x+8)/8)+x, int((p.y+16)/8))+y, 0)
		c := g.Fget(g.Mget(int((p.x+14)/8)+x, int((p.y+16)/8))+y, 0)
		shift := 0.0
		if !a && c {
			shift = -1
		}
		if !c && a {
			shift = 1
		}
		return (a || b || c), shift, 0
	}

	if facing == facingLeft {
		a := g.Fget(g.Mget(int((p.x+2)/8)+x, int((p.y+2)/8))+y, 0)
		b := g.Fget(g.Mget(int((p.x+2)/8)+x, int((p.y+8)/8))+y, 0)
		c := g.Fget(g.Mget(int((p.x+2)/8)+x, int((p.y+16)/8))+y, 0)
		shift := 0.0
		if !a && c {
			shift = -1
		}
		if !c && a {
			shift = 1
		}
		return (a || b || c), 0, shift
	}

	if facing == facingRight {
		a := g.Fget(g.Mget(int((p.x+14)/8)+x, int((p.y+2)/8))+y, 0)
		b := g.Fget(g.Mget(int((p.x+14)/8)+x, int((p.y+8)/8))+y, 0)
		c := g.Fget(g.Mget(int((p.x+14)/8)+x, int((p.y+16)/8))+y, 0)
		shift := 0.0
		if !a && c {
			shift = -1
		}
		if !c && a {
			shift = 1
		}
		return (a || b || c), 0, shift
	}

	return false, 0, 0
}

func sceanCollide(player vec2) bool {
	for _, c := range mainScean.collide {
		if c.collide(player) {
			return true
		}
	}
	return false
}
