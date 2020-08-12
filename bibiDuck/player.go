package main

import (
	"fantasyConsole/golf"
	"math"
)

var duck *player

type player struct {
	*sprite
	dx, dy                             float64
	maxDX, maxDY                       float64
	acc                                float64
	boost                              float64
	running, jumping, falling, sliding bool
	faceRight                          bool
	coyoteTime                         int
	hp                                 int
}

const coyoteTimeMax = 30

// Player animation states
const (
	waitRight = iota
	waitLeft
	runRight
	runLeft
	jumpRight
	jumpLeft
	fallRight
	fallLeft
	slideRight
	slideLeft
)

// Creates all the player animations and saves it to the duck variable
func initPlayer() {
	playerOpt := golf.SOp{TCol: golf.Col5, W: 2, H: 2}
	FplayerOpt := golf.SOp{TCol: golf.Col5, W: 2, H: 2, FH: true}
	sprite := newMob(40, 120, []drawable{
		&ani{ // Wait Right
			frames: []int{1, 3},
			speed:  0.05,
			o:      playerOpt,
		},
		&ani{ // Wait Left
			frames: []int{1, 3},
			speed:  0.05,
			o:      FplayerOpt,
		},
		&ani{ // Run Right
			frames: []int{5, 3, 7, 3},
			speed:  0.2,
			o:      playerOpt,
		},
		&ani{ // Run Left
			frames: []int{5, 3, 7, 3},
			speed:  0.2,
			o:      FplayerOpt,
		},
		&spr{ // Jump Up Right
			n: 9,
			o: playerOpt,
		},
		&spr{ // Jump Up Left
			n: 9,
			o: FplayerOpt,
		},
		&spr{ // Fall Down Right
			n: 11,
			o: playerOpt,
		},
		&spr{ // Fall Down Left
			n: 11,
			o: FplayerOpt,
		},
		&spr{ // Slide Right
			n: 13,
			o: golf.SOp{TCol: golf.Col5, W: 3, H: 2},
		},
		&spr{ // Slide Left
			n: 13,
			o: golf.SOp{TCol: golf.Col5, W: 3, H: 2, FH: true},
		},
	})

	duck = &player{
		sprite: sprite,
		maxDX:  2,
		maxDY:  3,
		acc:    0.5,
		boost:  5,
		hp:     3,
	}
}

// monitors controlls to move the player
func updatePlayer() {
	if duck.y > 192 {
		duck.hp--
		duck.x = 40
		duck.y = 120
	}
	duck.dy += gravity
	duck.dx *= friction
	if duck.dy > duck.maxDY {
		duck.dy = duck.maxDY
	}
	if duck.dx > duck.maxDX {
		duck.dx = duck.maxDX
	}

	// Run
	if g.Btn(golf.LeftArrow) {
		duck.dx -= duck.acc
		duck.running = true
		duck.faceRight = false
	}
	if g.Btn(golf.RightArrow) {
		duck.dx += duck.acc
		duck.running = true
		duck.faceRight = true
	}

	// Slide
	if duck.running &&
		!duck.falling &&
		!duck.jumping &&
		!g.Btn(golf.LeftArrow) &&
		!g.Btn(golf.RightArrow) {
		duck.running = false
		duck.sliding = true
	}

	// Jump
	if g.Btnp(golf.ZKey) &&
		duck.coyoteTime > 0 {
		duck.dy = 0
		duck.dy -= duck.boost
		duck.coyoteTime = 0
	}

	duck.x += duck.dx
	duck.y += duck.dy
	duck.coyoteTime--

	// Falling
	if duck.dy > 0 {
		duck.falling = true
		duck.jumping = false
	}

	// Jumping
	if duck.dy < 0.0 {
		duck.jumping = true
		duck.falling = false
	}

	// Check down collisions
	if duck.falling && (collideMap(duck.x+4, duck.y+14) || collideMap(duck.x+10, duck.y+14)) {
		duck.coyoteTime = coyoteTimeMax
		duck.falling = false
		duck.y -= duck.dy
		duck.dy = 0
	}

	// Check Left and Right collisions
	if collideMap(duck.x+4, duck.y+14) || collideMap(duck.x+10, duck.y+14) ||
		collideMap(duck.x, duck.y) || collideMap(duck.x+16, duck.y) {
		duck.x -= duck.dx
		duck.dx = 0
	}

	// Stop Sliding
	if duck.sliding {
		if math.Abs(duck.dx) < 0.2 || duck.running {
			duck.dx = 0
			duck.sliding = false
		}
	}
}

func drawPlayer() {
	// Animations depending on the duck state
	duck.a.(*stateAni).state = waitLeft
	if duck.faceRight {
		duck.a.(*stateAni).state = waitRight
	}

	if duck.running {
		duck.a.(*stateAni).state = runLeft
		if duck.faceRight {
			duck.a.(*stateAni).state = runRight
		}
	}

	if duck.sliding {
		duck.a.(*stateAni).state = slideLeft
		if duck.faceRight {
			duck.a.(*stateAni).state = slideRight
		}
	}

	if duck.jumping {
		duck.a.(*stateAni).state = jumpLeft
		if duck.faceRight {
			duck.a.(*stateAni).state = jumpRight
		}
	}

	if duck.falling {
		duck.a.(*stateAni).state = fallLeft
		if duck.faceRight {
			duck.a.(*stateAni).state = fallRight
		}
	}
}

func collideMap(x, y float64) bool {
	// Check collisions based on map tiles
	tile := g.Mget(int(x/8), int(y/8))
	return g.Fget(tile, 0)
}
