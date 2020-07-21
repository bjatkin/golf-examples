package main

import (
	"fantasyConsole/golf"
	"fmt"
)

var g *golf.Engine

func main() {
	g = golf.NewEngine(update, draw)

	g.BG(golf.Col0)
	g.LoadSprs(spriteSheet)
	g.LoadMap(mapData)
	initGame()
	g.Run()
}

var player int
var debug string

func initGame() {
	player = addEntity(entity{
		flags:   playerControlled,
		sig:     pos | spr | hp | collide,
		hp:      100,
		pos:     posComp{10, 10},
		spr:     sprComp{ani: [10]int{2, 3}, aniLen: 2, aniSpeed: 60, opt: golf.SOp{H: 2, TCol: golf.Col7}},
		collide: collidableComp{width: 8, height: 16, oldX: 10, oldY: 10},
	})

	//Wall
	addEntity(entity{
		sig:     pos | spr | collide,
		pos:     posComp{50, 150},
		spr:     sprComp{ani: [10]int{68}, aniLen: 1, opt: golf.SOp{H: 2, W: 4}},
		collide: collidableComp{width: 32, height: 16, oldX: 50, oldY: 150},
	})

	// for i := 0; i < 15; i++ {
	// addEnemy(rand.Intn(192), rand.Intn(192))
	// addEnemy(rand.Intn(192), rand.Intn(192))
	// }

	allUpdateSystems[movePlayer] = toSystem(playerControlled, pos, func(e *entity) {
		e.spr.ani = [10]int{2, 3}
		e.spr.aniLen = 2
		e.spr.aniSpeed = 60
		e.spr.opt.W = 1
		e.collide.width = 8
		if g.Btn(golf.WKey) {
			e.pos.y--
			e.spr.ani = [10]int{14, 16, 18, 20, -14, -16, -18, -20}
			e.spr.aniLen = 8
			e.spr.aniSpeed = 10
			e.spr.opt.W = 2
			e.collide.width = 16
		}
		if g.Btn(golf.SKey) {
			e.pos.y++
			e.spr.ani = [10]int{14, 16, 18, 20, -14, -16, -18, -20}
			e.spr.aniLen = 8
			e.spr.aniSpeed = 10
			e.spr.opt.W = 2
			e.collide.width = 16
		}
		if g.Btn(golf.AKey) {
			e.pos.x--
			e.spr.ani = [10]int{6, 8, 10, 12}
			e.spr.aniLen = 4
			e.spr.aniSpeed = 10
			e.spr.opt.FH = false
			e.spr.opt.W = 2
			e.collide.width = 16
		}
		if g.Btn(golf.DKey) {
			e.pos.x++
			e.spr.ani = [10]int{6, 8, 10, 12}
			e.spr.aniLen = 4
			e.spr.aniSpeed = 10
			e.spr.opt.FH = true
			e.spr.opt.W = 2
			e.collide.width = 16
		}
	})

	allUpdateSystems[doAI] = toSystem(0, ai|pos, func(e *entity) {
		t := allEntities[e.ai.target]
		dist := (((t.pos.x - e.pos.x) * (t.pos.x - e.pos.x)) + ((t.pos.y - e.pos.y) * (t.pos.y - e.pos.y)))
		if dist > e.ai.atkRange*e.ai.atkRange {
			return
		}
		e.spr.ani = [10]int{24, 25, 26}
		e.spr.aniLen = 3
		if e.pos.x < t.pos.x {
			e.pos.x++
			e.spr.opt.FH = true
		}
		if e.pos.x > t.pos.x {
			e.pos.x--
			e.spr.opt.FH = false
		}
		if e.pos.y < t.pos.y {
			e.pos.y++
		}
		if e.pos.y > t.pos.y {
			e.pos.y--
		}
	})

	allUpdateSystems[doAttack] = toSystem(enemy, ai|pos, func(e *entity) {
		t := &allEntities[e.ai.target]
		hDist := t.pos.x - e.pos.x
		vDist := t.pos.y - e.pos.y
		if hDist < 0 {
			hDist *= -1
		}
		if vDist < 0 {
			vDist *= -1
		}
		if hDist < 3 && vDist < 3 {
			t.hp--
		}
	})

	allUpdateSystems[startCollision] = toSystem(0, pos|collide, func(e *entity) {
		// Add all these entities to a list to make collision detection faster
		allCollidables[collidablePointer] = e
		collidablePointer++
	})

	allUpdateSystems[doCollision] = toSystem(0, pos|collide, doPhysics)

	allUpdateSystems[resolveCollision] = toSystem(0, pos|collide, func(e *entity) {
		e.pos.x = e.collide.oldX + e.collide.deltaX
		e.pos.y = e.collide.oldY + e.collide.deltaY
		e.collide.oldX = e.pos.x
		e.collide.oldY = e.pos.y
		collidablePointer = 0
	})

	allDrawSystems[drawSprite] = toSystem(0, pos|spr, func(e *entity) {
		e.spr.frame++
		if e.spr.aniSpeed == 0 {
			e.spr.aniSpeed = 1
		}
		if e.spr.frame%e.spr.aniSpeed == 0 {
			e.spr.aniFrame++
		}
		if e.spr.aniFrame >= e.spr.aniLen {
			e.spr.aniFrame = 0
		}
		frame := e.spr.ani[e.spr.aniFrame]
		opt := e.spr.opt
		if frame < 0 {
			opt.FH = true
			frame *= -1
		}
		g.Spr(frame, e.pos.x, e.pos.y, opt)
		if e.hasComponent(collide) {
			g.Rect(e.pos.x, e.pos.y, e.collide.width, e.collide.height, golf.Col1)
		}
	})

	allDrawSystems[drawHP] = toSystem(playerControlled, hp, func(e *entity) {
		if e.hp < 0 {
			g.TextL("You DED!", golf.TOp{Col: golf.Col3})
			return
		}
		g.TextL(fmt.Sprintf("HP: %d", e.hp), golf.TOp{Col: golf.Col3})
	})
}

func update() {
	runUpdateSystems()
}

func draw() {
	g.Cls()
	runDrawSystems()
	a, b := allCollidables[0], allCollidables[1]
	x1, y1 := a.collide.oldX+a.collide.deltaX, a.collide.oldY+a.collide.deltaY
	w1, h1 := a.collide.width, a.collide.height
	x2, y2 := b.collide.oldX+b.collide.deltaX, b.collide.oldY+b.collide.deltaY
	w2, h2 := b.collide.width, b.collide.height
	debug += fmt.Sprintf("\nx1 %d, y1 %d, w1 %d, h1 %d\nx2 %d, y2 %d, w1 %d, h1 %d", int(x1), int(y1), int(w1), int(h1), int(x2), int(y2), int(w2), int(h2))
	var xO1, yO1, xO2, yO2, xO3, yO3 bool
	if x1 >= x2 && x1 <= x2+h2 {
		xO1 = true
	}
	if x1+h1 >= x2 && x1+h1 <= x2+h2 {
		xO2 = true
	}
	if x1 >= x2 && x1+h1 <= x2+h2 {
		xO3 = true
	}
	if y1 >= y2 && y1 <= y2+h2 {
		yO1 = true
	}
	if y1+h1 >= y2 && y1+h1 <= y2+h2 {
		yO2 = true
	}
	if y1 >= y2 && y1+h1 <= y2+h2 {
		yO3 = true
	}
	debug += fmt.Sprintf("\nxO1: %v, yO1: %v,\nxO2: %v, yO2: %v\nxO3: %v, yO3: %v", xO1, yO1, xO2, yO2, xO3, yO3)

	g.TextR(debug, golf.TOp{Col: golf.Col3})
	debug = ""
}

func addEnemy(x, y float64) int {
	// zombie
	id := addEntity(entity{
		flags:   enemy,
		sig:     pos | spr | ai | collide,
		pos:     posComp{x, y},
		spr:     sprComp{ani: [10]int{24, 25, 26}, aniLen: 3, aniSpeed: 30, opt: golf.SOp{H: 2, TCol: golf.Col7}},
		ai:      aiComp{atkRange: 150, target: player},
		collide: collidableComp{width: 8, height: 16, oldX: x, oldY: y},
	})
	return id
}
