package main

import (
	"fantasyConsole/golf"
	"math"
	"math/rand"
	"strconv"
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
	hp                                 int
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
	initGopher()
	initHUD()
	initConvo()
	initParticles()

	g.Run()
}

var drawMap bool
var cameraX, cameraY int

func update() {
	if !gopherConvo.running {
		updatePlayer()
	}
	updateCamera()
	checkPickups()

	if g.Btnp(golf.XKey) && nearGopher {
		if duck.s.x < 387 {
			duck.s.a.(*stateAni).state = waitRight
		} else {
			duck.s.a.(*stateAni).state = waitLeft
		}
		gopherConvo.next()
	}
	g.Camera(cameraX, cameraY)
}

func draw() {
	g.Cls(golf.Col5)
	g.Map(0, 0, 128, 128, 0, 0, alpha)

	if !gopherConvo.running {
		drawPlayer()
	}
	drawSprites()
	drawHUD()
	gopherConvo.draw()
	drawSpeachBubble()
	drawParticles()
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
		hp:    3,
	}
}

func initGopher() {
	for x := 0; x < 128; x++ {
		for y := 0; y < 128; y++ {
			if g.Mget(x, y) == 201 {
				g.Mset(x, y, 0)
				g.Mset(x+1, y, 0)
				g.Mset(x, y+1, 0)
				g.Mset(x+1, y+1, 0)
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

var nearGopher bool

func drawSpeachBubble() {
	x, y := 387.0, 125.0
	dx := math.Abs(duck.s.x - x)
	dy := math.Abs(duck.s.y - y)
	if dx > 30 || dy > 30 {
		nearGopher = false
		return
	}
	nearGopher = true
	g.Spr(73, x, y, golf.SOp{W: 2, H: 2, TCol: golf.Col5})
	yy := y + 3
	if (g.Frames()/30)%2 == 0 {
		yy = y + 4
	}
	g.Text(x+4, y+4, "(x)", golf.TOp{Col: golf.Col4})
	g.Text(x+6, y+4, "(x)", golf.TOp{Col: golf.Col4})
	g.Text(x+5, y+4, "(x)", golf.TOp{Col: golf.Col4})
	g.Text(x+4, yy, "(x)")
	g.Text(x+6, yy, "(x)")
	g.Text(x+5, yy, "(x)")
}

func updatePlayer() {
	if duck.s.y > 192 {
		duck.hp--
		duck.s.x = 40
		duck.s.y = 120
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
		if math.Abs(duck.dx) < 0.2 || duck.running {
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
			collectedFethers++
			fether.delete()
			for p := 0; p < 20; p++ {
				addParticle(fether.x, fether.y+10)
			}
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
			collectedEggs++
			egg.a.(*stateAni).state++
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

var shinyFether *ani
var shinyEgg *ani
var collectedFethers int
var collectedEggs int

func initHUD() {
	shinyFether = &ani{
		frames: []int{32, 32, 32, 32, 32, 32, 32, 32, 192, 224, 256, 288},
		speed:  0.1,
		o:      golf.SOp{TCol: golf.Col5, Fixed: true},
	}

	shinyEgg = &ani{
		frames: []int{193, 193, 193, 193, 193, 193, 193, 193, 257, 259, 261, 263},
		speed:  0.1,
		o:      golf.SOp{W: 2, H: 2, TCol: golf.Col5, Fixed: true},
	}
}

func drawHUD() {
	fether := 95.0
	eggs := 140.0

	// Outline
	g.RectFill(0, 172, 192, 20, golf.Col7, true)
	g.Rect(0, 172, 192, 20, golf.Col0, true)

	// HP
	hearts := "HP "
	for i := 0; i < 3; i++ {
		if i < duck.hp {
			hearts += "<3"
		} else {
			hearts += "<4"
		}
	}
	g.Text(8, 177, hearts, golf.TOp{SW: 2, SH: 2, Fixed: true})

	// Fether
	g.RectFill(fether-5, 174, 38, 16, golf.Col6, true)
	shinyFether.draw(fether, 179)
	g.Spr(16, fether+8, 179, golf.SOp{TCol: golf.Col5, Fixed: true})
	g.Text(fether+16, 180, strconv.Itoa(collectedFethers), golf.TOp{Fixed: true})

	// Egg
	g.RectFill(eggs-5, 174, 44, 16, golf.Col6, true)
	shinyEgg.draw(eggs, 174)
	g.Spr(16, eggs+14, 179, golf.SOp{TCol: golf.Col5, Fixed: true})
	g.Text(eggs+22, 180, strconv.Itoa(collectedEggs), golf.TOp{Fixed: true})
}

type convo struct {
	portrait   []int
	lines      []string
	line       int
	height     int
	goalHeight int
	running    bool
}

func (c *convo) draw() {
	g.RectFill(0, float64(c.height), 192, 34, golf.Col7, true)
	g.Rect(0, float64(c.height), 192, 34, golf.Col0, true)
	g.Spr(c.portrait[c.line], 1, float64(c.height)+2, golf.SOp{W: 4, H: 4, TCol: golf.Col5, Fixed: true})
	if c.height < c.goalHeight {
		c.height++
	}
	if c.height > c.goalHeight {
		c.height--
	}
	g.Text(35, float64(c.height+4), c.lines[c.line], golf.TOp{Fixed: true})
}

func (c *convo) next() {
	if !c.running {
		c.running = true
		c.goalHeight = 158
		c.line = 0
	}
	if c.height != c.goalHeight {
		c.height = c.goalHeight
		return
	}
	c.line++
	if c.line >= len(c.lines) {
		c.line--
		c.goalHeight = 192
		c.running = false
	}
}

var gopherConvo *convo

func initConvo() {
	gopherConvo = &convo{
		portrait: []int{65, 69, 65},
		lines: []string{
			"BIBI DUCK: Hey, do you\nknow how to get out of\nthis place!?",
			"JOE GOPHER: Why would you\nwant to leave?\nThis place is great!",
			"BIBI DUCK: Oh...",
		},
		height:     192,
		goalHeight: 192,
	}
}

type particle struct {
	x, y float64
	life int
}

var allParticles = [100]*particle{}
var pPointer = 0

func initParticles() {
	for i := 0; i < len(allParticles); i++ {
		allParticles[i] = &particle{0, 0, 0}
	}
}

func addParticle(x, y float64) {
	allParticles[pPointer] = &particle{x, y, 20}
	pPointer++
	if pPointer >= len(allParticles) {
		pPointer = 0
	}
}

func drawParticles() {
	for _, p := range allParticles {
		if p.life <= 0 {
			continue
		}
		p.x += float64(rand.Intn(3) - 1)
		p.y += float64(rand.Intn(5) - 3)
		p.life--
		g.Pset(p.x-float64(cameraX), p.y, golf.Col7)
	}
}
