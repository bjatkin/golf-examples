package main

import (
	"fantasyConsole/golf"
	"math"
)

var g *golf.Engine
var alpha = golf.SOp{TCol: golf.Col5}

type player struct {
	s                                  *sprite
	dx, dy                             float64
	maxDX, maxDY                       float64
	acc                                float64
	boost                              float64
	running, jumping, falling, sliding bool
	faceRight                          bool
	coyoteTime                         int
}

var duck *player

const gravity = 0.3
const friction = 0.85
const coyoteTimeMax = 30

func main() {
	g = golf.NewEngine(update, draw)

	g.LoadMap(mapData)
	g.LoadSprs(spriteSheet)
	g.LoadFlags(spriteFlags)

	g.PalA(golf.Pal7)
	g.PalB(golf.Pal8)

	g.RAM[0x3601] = 0

	initPickups()
	initPlayer()

	g.Run()
}

var drawMap bool
var cameraX, cameraY int

func update() {
	updatePlayer()
	updateCamera()
	checkPickups()

	g.Camera(cameraX, cameraY)
}

func draw() {
	g.Cls(golf.Col5)

	g.Map(0, 0, 128, 128, 0, 0, alpha)

	drawPlayer()
	drawSprites()
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
		s:     sprite,
		maxDX: 2,
		maxDY: 3,
		acc:   0.5,
		boost: 5,
	}
}

func updatePlayer() {
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
	if g.Btnp(golf.Space) &&
		duck.coyoteTime > 0 {
		duck.dy = 0
		duck.dy -= duck.boost
		duck.coyoteTime = 0
	}

	duck.s.x += duck.dx
	duck.s.y += duck.dy
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

	// Check down
	if duck.falling && (collideMap(duck.s.x+4, duck.s.y+14) || collideMap(duck.s.x+10, duck.s.y+14)) {
		duck.coyoteTime = coyoteTimeMax
		duck.falling = false
		duck.s.y -= duck.dy
		duck.dy = 0
	}

	// Check Left and Right
	if collideMap(duck.s.x+4, duck.s.y+14) || collideMap(duck.s.x+10, duck.s.y+14) ||
		collideMap(duck.s.x, duck.s.y) || collideMap(duck.s.x+16, duck.s.y) {
		duck.s.x -= duck.dx
		duck.dx = 0
	}

	// Stop Sliding
	if duck.sliding {
		if math.Abs(duck.dx) < 0.75 || duck.running {
			duck.dx = 0
			duck.sliding = false
		}
	}
}

func drawPlayer() {
	// Animations
	duck.s.a.(*stateAni).state = waitLeft
	if duck.faceRight {
		duck.s.a.(*stateAni).state = waitRight
	}

	if duck.running {
		duck.s.a.(*stateAni).state = runLeft
		if duck.faceRight {
			duck.s.a.(*stateAni).state = runRight
		}
	}

	if duck.sliding {
		duck.s.a.(*stateAni).state = slideLeft
		if duck.faceRight {
			duck.s.a.(*stateAni).state = slideRight
		}
	}

	if duck.jumping {
		duck.s.a.(*stateAni).state = jumpLeft
		if duck.faceRight {
			duck.s.a.(*stateAni).state = jumpRight
		}
	}

	if duck.falling {
		duck.s.a.(*stateAni).state = fallLeft
		if duck.faceRight {
			duck.s.a.(*stateAni).state = fallRight
		}
	}
}

func updateCamera() {
	cameraX = int(duck.s.x) - 96
	if cameraX < 0 {
		cameraX = 0
	}
	if cameraX > 422 {
		cameraX = 422
	}
}

func checkPickups() {
	remove := -1
	for i, fether := range allFethers {
		dx := math.Abs(fether.x - duck.s.x)
		dy := math.Abs(fether.y - duck.s.y)
		if dx < 16 && dy < 16 {
			fether.delete()
			remove = i
			break
		}
	}
	if remove != -1 {
		allFethers[remove] = allFethers[len(allFethers)-1]
		allFethers = allFethers[:len(allFethers)-1]
	}

	remove = -1
	for i, egg := range allEggs {
		dx := math.Abs(egg.x - duck.s.x)
		dy := math.Abs(egg.y - duck.s.y)
		if dx < 16 && dy < 16 {
			egg.delete()
			remove = i
			break
		}
	}
	if remove != -1 {
		allEggs[remove] = allEggs[len(allEggs)-1]
		allEggs = allEggs[:len(allEggs)-1]
	}
}

func collideMap(x, y float64) bool {
	tile := g.Mget(int(x/8), int(y/8))
	return g.Fget(tile, 0)
}
